package controllers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/thetkpark/golang-todo/data"
	"github.com/thetkpark/golang-todo/services"
	"gorm.io/gorm"
)

type Controller struct {
	db         *gorm.DB
	jwtManager *services.JWTManager
}

func NewController(db *gorm.DB, jwtManager *services.JWTManager) *Controller {
	return &Controller{db: db, jwtManager: jwtManager}
}

type AuthController struct {
	userRepository data.UserRepository
	jwtManager     *services.JWTManager
	log            hclog.Logger
}

func NewAuthController(userRepo data.UserRepository, jwtManager *services.JWTManager, log hclog.Logger) *AuthController {
	return &AuthController{
		userRepository: userRepo,
		jwtManager:     jwtManager,
		log:            log,
	}
}

type TodoController struct {
	todoRepository data.TodoRepository
	log            hclog.Logger
}

func NewTodoController(todoRepo data.TodoRepository, log hclog.Logger) *TodoController {
	return &TodoController{
		todoRepository: todoRepo,
		log:            log,
	}
}
