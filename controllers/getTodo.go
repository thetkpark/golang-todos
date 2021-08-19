package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
)

// GetTodoController getTodo
// @Summary Get all todos
// @Tags todo
// @Produce  json
// @Security JwtAuth
// @Success 200 {object} models.Todo "the list todo that user have"
// @Failure 401 {object} controllers.ErrorMessage "Unauthorized"
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/todo [get]
func (c *Controller) GetTodoController(ctx *gin.Context) {
	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	var todos []models.Todo
	if tx := c.db.Where(&models.Todo{UserId: userId}).Find(&todos); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
	}

	ctx.JSON(200, todos)
}
