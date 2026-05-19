package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:80/api/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()

	done := make(chan bool)

	go func() {
		// Read message from the server.
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			fmt.Printf("Received: %s\n", message)
		}

		done <- true
	}()

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}

			if text == "" {
				break
			}

			text = strings.TrimSpace(text)

			err = ws.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}()

	<-done
}
