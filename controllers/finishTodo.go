package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
	"gorm.io/gorm"
)

func (c *Controller) FinishTodoController(ctx *gin.Context) {
	todoId := ctx.Param("todoId")

	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	var todo models.Todo
	tx := c.db.Where(`id = ? AND user_id = ?`, todoId, userId).First(&todo)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{
				"message": "Not Found",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
		return
	}

	todo.IsFinished = true
	tx = c.db.Save(&todo)
	if tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
		return
	}

	var todos []models.Todo
	if tx := c.db.Where(&models.Todo{UserId: userId}).Find(&todos); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
	}

	ctx.JSON(200, todos)
}
