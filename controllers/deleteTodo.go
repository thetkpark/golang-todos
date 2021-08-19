package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
	"gorm.io/gorm"
)

// DeleteTodoController deleteTodo
// @Summary Delete todo
// @Tags todo
// @Accept json
// @Produce  json
// @Security JwtAuth
// @Param todoId path integer true "id of todo to delete"
// @Success 200 {array} models.Todo "the list of todos that user have"
// @Failure 401
// @Failure 404 {object} controllers.ErrorMessage "Todo not found"
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/todo/{todoId} [delete]
func (c *Controller) DeleteTodoController(ctx *gin.Context) {
	todoId := ctx.Param("todoId")

	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	var todo models.Todo
	tx := c.db.Where(`id = ? AND user_id = ?`, todoId, userId).First(&todo)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(404, ErrorMessage{fmt.Sprintf("todo with id %s id not found", todoId)})
			return
		}
		ctx.JSON(500, ErrorMessage{tx.Error.Error()})
		return
	}

	tx = c.db.Delete(&todo)
	if tx.Error != nil {
		ctx.JSON(500, ErrorMessage{tx.Error.Error()})
		return
	}

	var todos []models.Todo
	if tx := c.db.Where(&models.Todo{UserId: userId}).Find(&todos); tx.Error != nil {
		ctx.JSON(500, ErrorMessage{tx.Error.Error()})
	}

	ctx.JSON(200, todos)
}
