package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// websocket client:
// connect
// read event handler
// write event handler
// broadcast event handler
type broadcastData struct {
	Data int
}

// ClientManager websocket client Manager struct
// 管理三個 channel。
// register, unregister 記錄連線
// broadcast 管理 廣播
type clientManager struct {
	clients    []*wsClient
	register   chan *wsClient
	unRegister chan *wsClient
	broadcast  chan *broadcastData
}

// WsManager WebSocket 管理器
// 三個通道
var WsManager = clientManager{
	clients:    make([]*wsClient, 100),
	register:   make(chan *wsClient),
	unRegister: make(chan *wsClient),
	broadcast:  make(chan *broadcastData, 10),
}

func (manager *clientManager) RegisterClient(ctx *gin.Context) {
	// 升級連線 -> 註冊 client (goruntime) -> read, write goruntime
	upgrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		println("websocket client connect %v error", ctx.Param("channel"))
		return
	}

	client := &wsClient{
		ID:     uuid.NewV4().String(),
		Socket: conn,
		Send:   make(chan int),
	}

	// 註冊了 client
	manager.register <- client

	go client.Read()
	go client.Write()
}

func (manager *clientManager) Broadcast(message int) {
	data := &broadcastData{
		Data: message,
	}
	manager.broadcast <- data
}

func (manager *clientManager) Start() {
	count := 0

	println("Websocket manage start")
	for {
		select {
		case client := <-manager.register:
			println("Websocket client %s connect", client.ID)
			count++
			manager.Broadcast(count)
		case client := <-manager.unRegister:
			println("Unregister websocket client %s", client.ID)

			close(client.Send)

		case data := <-manager.broadcast:

			for _, conn := range manager.clients {
				conn.Send <- data.Data
			}
		}
	}
}
