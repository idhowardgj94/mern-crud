package routes

import (
	"github.com/gin-gonic/gin"
)

// UserRoute user's route
func UserRoute(users *gin.RouterGroup) {
	users.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"debug":  c.Param("id"),
		})
	})
	users.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	users.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	users.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"debug":  c.Param("id"),
		})
	})
	users.DELETE("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"degbug": c.Param("id"),
		})
	})
}
