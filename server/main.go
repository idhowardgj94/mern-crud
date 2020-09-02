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
	println("start")
	go websocket.WsManager.Start()
	// webscoket
	r.GET("/test", func(c *gin.Context) {
		websocket.WsManager.RegisterClient(c)
		println("到底在做什麼小叮噹拉")
	})
	println("done")
	r.Run()
}
