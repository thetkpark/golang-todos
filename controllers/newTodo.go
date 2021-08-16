package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
)

type NewTodoDto struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}

func (c *Controller) CreateTodoController(ctx *gin.Context) {
	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	var body NewTodoDto
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	todo := models.Todo{
		Title:      body.Title,
		UserId:     userId,
		IsFinished: false,
	}

	if tx := c.db.Create(&todo); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
	}

	var todos []models.Todo
	if tx := c.db.Where(&models.Todo{UserId: userId}).Find(&todos); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
	}

	ctx.JSON(201, todos)
}
