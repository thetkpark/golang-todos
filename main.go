package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thetkpark/golang-todo/controllers"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/middlewares"
	"github.com/thetkpark/golang-todo/models"
	"log"
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
	err = db.AutoMigrate(&models.User{}, &models.Todo{})
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
	router.POST("/api/signin", controllers.SignInController)

	authorization := router.Group("/")
	authorization.Use(middlewares.AuthorizeJWT())
	{
		authorization.POST("/api/todo", controllers.CreateTodoController)
		authorization.PATCH("/api/todo/:todoId", controllers.FinishTodoController)
		authorization.DELETE("/api/todo/:todoId", controllers.DeleteTodoController)
	}

	err = router.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}
}
