package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thetkpark/golang-todo/controllers"
	"github.com/thetkpark/golang-todo/data"
	_ "github.com/thetkpark/golang-todo/docs"
	"github.com/thetkpark/golang-todo/middlewares"
	"github.com/thetkpark/golang-todo/models"
	"github.com/thetkpark/golang-todo/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Swagger Golang Todo API
// @version 1.0
// @description This is a sample of API server that store todos
//
// @license.name MIT
//
// @host localhost:5000
// @BasePath /
//
// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization

func main() {

	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalln("error loading .env")
		}
	}

	logger := hclog.Default()
	router := gin.Default()

	gormDB, err := data.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = gormDB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}

	// Create JWTManager
	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) == 0 {
		log.Fatalln("cannot get JWT_SECRET from OS env")
	}
	jwtManager := services.NewJWTManager(jwtSecret, time.Hour*24)

	// Create data repository
	userRepository := data.NewGormUserRepository(gormDB, logger)
	todoRepository := data.NewGormTodoRepository(gormDB, logger)

	// Create controller
	authController := controllers.NewAuthController(userRepository, jwtManager, logger)
	todoController := controllers.NewTodoController(todoRepository, logger)

	// Create middleware
	middleware := middlewares.NewMiddleware(jwtManager)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			`success`:   true,
			`timestamp`: time.Now(),
		})
	})
	router.POST("/api/regis", authController.RegisterController)
	router.POST("/api/signin", authController.SignInController)

	authorization := router.Group("/")
	authorization.Use(middleware.AuthorizeJWT())
	{
		authorization.GET("/api/todo", todoController.GetTodoController)
		authorization.POST("/api/todo", todoController.CreateTodoController)
		authorization.PATCH("/api/todo/:todoId", todoController.FinishTodoController)
		authorization.DELETE("/api/todo/:todoId", todoController.DeleteTodoController)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	// Initializing the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Block until quit channel is received
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
