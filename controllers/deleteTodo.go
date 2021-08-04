package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
	"gorm.io/gorm"
)

func DeleteTodoController(ctx *gin.Context) {
	todoId := ctx.Param("todoId")

	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	db, err := db.GetDB()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	var todo models.Todo
	tx := db.Where(`id = ? AND user_id = ?`, todoId, userId).First(&todo)
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

	tx = db.Delete(&todo)
	if tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
		return
	}

	var todos []models.Todo
	if tx := db.Where(&models.Todo{UserId: userId}).Find(&todos); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
	}

	ctx.JSON(200, todos)
}
