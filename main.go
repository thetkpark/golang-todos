package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			`success`: true,
			`timestamp`: time.Now(),
		})
	})

	err := router.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}
}
