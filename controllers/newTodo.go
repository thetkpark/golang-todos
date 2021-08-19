package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
)

type NewTodoDto struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}

// CreateTodoController createTodo
// @Summary Create todo
// @Tags todo
// @Accept json
// @Produce  json
// @Security JwtAuth
// @Param body body controllers.NewTodoDto true "title of todo to create"
// @Success 201 {array} models.Todo "the list of todos that user have"
// @Failure 401
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/todo [post]
func (c *Controller) CreateTodoController(ctx *gin.Context) {
	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	var body NewTodoDto
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(400, ErrorMessage{err.Error()})
		return
	}

	todo := models.Todo{
		Title:      body.Title,
		UserId:     userId,
		IsFinished: false,
	}

	if tx := c.db.Create(&todo); tx.Error != nil {
		ctx.JSON(500, ErrorMessage{tx.Error.Error()})
	}

	var todos []models.Todo
	if tx := c.db.Where(&models.Todo{UserId: userId}).Find(&todos); tx.Error != nil {
		ctx.JSON(500, ErrorMessage{tx.Error.Error()})
	}

	ctx.JSON(201, todos)
}
