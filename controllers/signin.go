package controllers

import "github.com/gin-gonic/gin"

type SigninDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

func SignInController(ctx *gin.Context) {
	var body SigninDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}
