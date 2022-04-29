package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"message"`
	IsCompleted bool   `json:"is_completed"`
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	r := gin.Default()
	r.GET("/", Ping)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
