package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	staticDir     string
	last_id       uint
	value         int
	clients       map[*ClientHandler]uint
	broadcastChan chan string
}

func NewServer(staticDirPath string) *Server {
	var clients = make(map[*ClientHandler]uint)
	var broadcastChan = make(chan string, 64)

	return &Server{
		staticDir:     staticDirPath,
		last_id:       0,
		value:         42,
		clients:       clients,
		broadcastChan: broadcastChan,
	}
}

func (s *Server) Run() error {
	routes := s.Routes()
	fmt.Println("Listening on: localhost:80")
	go s.broadcast()
	go s.updateLoop()
	return http.ListenAndServe(":80", routes)
}

func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Build and register the API layer
	api := s.apiRoutes()
	mux.Handle("/api/", http.StripPrefix("/api", api))

	// Build file server
	fs := http.FileServer(http.Dir(s.staticDir))
	mux.HandleFunc("/", s.spaHandler(fs))

	return mux
}

func (s *Server) apiRoutes() *http.ServeMux {
	api := http.NewServeMux()

	api.HandleFunc("/ws", s.serveWs)
	api.HandleFunc("/health", s.HealthHandler)

	return api
}

type healthResponse struct {
	Message string `json:"message"`
}

func (s *Server) HealthHandler(w http.ResponseWriter, _ *http.Request) {
	message := "Hello from json API"
	response := healthResponse{Message: message}

	json.NewEncoder(w).Encode(response)
}

func (s *Server) spaHandler(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// the route should have caught this already so not found
		if strings.HasPrefix(path, "/api") {
			http.NotFound(w, r)
			return
		}

		// route through spa if theres no file to match
		_, err := os.Stat(filepath.Join(s.staticDir, path))
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(s.staticDir, "index.html"))
			return
		}

		// give a static file otherwise
		fs.ServeHTTP(w, r)
	}
}

type Update struct {
	Value int
}

func (s *Server) updateLoop() {
	ticker := time.NewTicker(5000 * time.Millisecond)

	for range ticker.C {
		update := Update{s.value}
		newMsg, err := json.Marshal(update)
		if err != nil {
			fmt.Println("bad update")
			continue
		}

		for client := range s.clients {
			err := client.ws.WriteMessage(websocket.TextMessage, newMsg)
			if err == nil {
				continue
			}

			fmt.Println("broadcast error:", client.id, err)
			client.ws.Close()
			delete(s.clients, client)
		}

	}
}

func (s *Server) broadcast() {
	for newMsg := range s.broadcastChan {
		var update Update

		err := json.Unmarshal([]byte(newMsg), &update)
		if err == nil && update.Value != 0 {
			fmt.Println("Got:", update)
			s.value = update.Value
			continue
		}

		fmt.Println(err)

		for client := range s.clients {
			err := client.ws.WriteMessage(websocket.TextMessage, []byte(newMsg))

			if err == nil {
				continue
			}

			fmt.Println("broadcast error:", client.id, err)
			client.ws.Close()
			delete(s.clients, client)
		}
	}
}

func (s *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("whoops")
		return
	}

	id := s.last_id
	s.last_id += 1

	handler := NewClientHandler(id, ws, s.broadcastChan)

	s.clients[handler] = id

	go handler.Run()
}
