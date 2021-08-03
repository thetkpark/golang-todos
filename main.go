package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/controllers"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	db, err := db.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&models.Users{}, &models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			`success`:   true,
			`timestamp`: time.Now(),
		})
	})

	router.POST("/api/regis", controllers.RegisterController)

	err = router.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}
}
