package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thetkpark/golang-todo/controllers"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/middlewares"
	"github.com/thetkpark/golang-todo/models"
	"log"
	"os"
	"time"
)

func main() {

	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalln("error loading .env")
		}
	}

	router := gin.Default()

	gormDB, err := db.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = gormDB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}

	// Create controller
	controller := controllers.NewController(gormDB)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			`success`:   true,
			`timestamp`: time.Now(),
		})
	})
	router.POST("/api/regis", controller.RegisterController)
	router.POST("/api/signin", controller.SignInController)

	authorization := router.Group("/")
	authorization.Use(middlewares.AuthorizeJWT())
	{
		authorization.GET("/api/todo", controller.GetTodoController)
		authorization.POST("/api/todo", controller.CreateTodoController)
		authorization.PATCH("/api/todo/:todoId", controller.FinishTodoController)
		authorization.DELETE("/api/todo/:todoId", controller.DeleteTodoController)
	}

	err = router.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Running on 5000")
}
