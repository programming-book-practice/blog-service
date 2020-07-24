package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//c.String(200, "Hello, Gin User")
		c.JSON(200, gin.H{"message", "pong"}) //./main.go:11:21: missing key in map literal
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
