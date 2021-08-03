package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

func RegisterController(ctx *gin.Context) {
	var bodyData RegisterDto
	if err := ctx.ShouldBindJSON(&bodyData); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bodyData.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.Users{
		Username: bodyData.Username,
		Password: string(hashedPassword),
	}

	db, err := db.GetDB()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tx := db.Create(&user); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(201, user)
}
