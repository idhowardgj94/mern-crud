package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// wsClient Websocket 客戶端
type wsClient struct {
	ID     string
	Socket *websocket.Conn
	Send   chan int
}

func (c *wsClient) Read() {
	defer func() {
		WsManager.unRegister <- c
		c.Socket.Close()
	}()

	for {
		_, _, err := c.Socket.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (c *wsClient) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", message)))
		}
	}
}
