package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thetkpark/golang-todo/controllers"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
	"log"
	"net/http"
	"time"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env")
	}

	router := gin.Default()

	db, err := db.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&models.Users{}, &models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(controllers.Error); ok {
			fmt.Printf("%T", err)
			c.JSON(int(err.StatusCode), gin.H{
				"message": err.Message,
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

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
