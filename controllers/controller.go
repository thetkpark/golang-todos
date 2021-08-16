package controllers

import (
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
