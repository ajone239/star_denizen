package internal

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type ClientHandler struct {
	id            uint
	ws            *websocket.Conn
	broadcastChan chan<- string
}

func NewClientHandler(
	id uint,
	ws *websocket.Conn,
	broadcastChan chan<- string,
) *ClientHandler {
	return &ClientHandler{
		id:            id,
		ws:            ws,
		broadcastChan: broadcastChan,
	}
}

func (c *ClientHandler) Run() {
	msg := fmt.Sprintf("You are %d", c.id)
	if err := c.ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		return
	}

	for {
		// Read message from browser
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		newMsg := string(msg)

		fmt.Println(newMsg)
		c.broadcastChan <- newMsg
	}
}
