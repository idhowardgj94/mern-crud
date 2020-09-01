package main

import (
	"mem-crud-go/routes"
	"mem-crud-go/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	// default route
	r := gin.Default()
	users := r.Group("/api/users")
	routes.UserRoute(users)
	websocket.WsManager.start()
	r.Run()
}
