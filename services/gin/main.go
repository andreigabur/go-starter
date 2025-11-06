package main

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", handlers.HandleUsers)
	router.Run() // listens on 0.0.0.0:8080 by default
}
