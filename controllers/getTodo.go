package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
)

func GetTodoController(ctx *gin.Context) {
	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	db, err := db.GetDB()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
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
