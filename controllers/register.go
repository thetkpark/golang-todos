package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
	"github.com/thetkpark/golang-todo/services"
	"golang.org/x/crypto/bcrypt"
)

type RegisterDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

func RegisterController(ctx *gin.Context) {
	var bodyData RegisterDto
	if err := ctx.ShouldBindJSON(&bodyData); err != nil {
		panic(Error{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bodyData.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(Error{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	user := models.Users{
		Username: bodyData.Username,
		Password: string(hashedPassword),
	}

	db, err := db.GetDB()
	if err != nil {
		panic(Error{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	if tx := db.Create(&user); tx.Error != nil {
		panic(Error{
			StatusCode: 500,
			Message:    tx.Error.Error(),
		})
	}

	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		panic(Error{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"token": token,
	})
}
