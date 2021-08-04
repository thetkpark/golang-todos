package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
)

type SigninDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

func SignInController(ctx *gin.Context) {
	var body SigninDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		panic(Error{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	db, err := db.GetDB()
	if err != nil {
		panic(Error{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	var user models.User
	db.Where(&models.User{Username: body.Username}).First(&user)
}
