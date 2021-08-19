package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetTodoController getTodo
// @Summary Get all todos
// @Tags todo
// @Produce  json
// @Security JwtAuth
// @Success 200 {array} models.Todo "the list of todos that user have"
// @Failure 401
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/todo [get]
func (c *TodoController) GetTodoController(ctx *gin.Context) {
	v, _ := ctx.Get("userId")
	var userId = uint(v.(float64))

	todos, err := c.todoRepository.FindAll(userId)
	if err != nil {
		ctx.JSON(500, ErrorMessage{err.Error()})
	}

	ctx.JSON(200, todos)
}
