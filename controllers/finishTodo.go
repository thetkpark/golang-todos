package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// FinishTodoController finishTodo
// @Summary Finish todo
// @Tags todo
// @Accept json
// @Produce  json
// @Security JwtAuth
// @Param todoId path integer true "id of todo that finished"
// @Success 200 {array} models.Todo "the list of todos that user have"
// @Failure 401
// @Failure 404 {object} controllers.ErrorMessage "Todo not found"
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/todo/{todoId} [patch]
func (c *TodoController) FinishTodoController(ctx *gin.Context) {
	todoId, err := strconv.Atoi(ctx.Param("todoId"))
	if err != nil {
		c.log.Error("cannot convert todoId string to int")
		ctx.JSON(500, ErrorMessage{err.Error()})
		return
	}

	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	_, err = c.todoRepository.FinishTodo(uint(todoId), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, ErrorMessage{fmt.Sprintf("todo with id %d is not found", todoId)})
			return
		}
		c.log.Error("error finishing todo", err.Error())
		ctx.JSON(500, ErrorMessage{err.Error()})
		return
	}

	todos, err := c.todoRepository.FindAll(userId)
	if err != nil {
		ctx.JSON(500, ErrorMessage{err.Error()})
		return
	}

	ctx.JSON(200, todos)
}
