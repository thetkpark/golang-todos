package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
)

type RegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterController(ctx *gin.Context) {
	var bodyData RegisterDto
	if err := ctx.BindJSON(&bodyData); err != nil {
		fmt.Println("Error Occure while paring ")
		ctx.JSON(500, err)
		return
	}

	user := models.Users{
		Username: bodyData.Username,
		Password: bodyData.Password,
	}

	db, err := db.GetDB()
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if tx := db.Create(&user); tx.Error != nil {
		ctx.JSON(500, tx.Error)
		return
	}

	ctx.JSON(201, user)
}
