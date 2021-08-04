package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/db"
	"github.com/thetkpark/golang-todo/models"
	"github.com/thetkpark/golang-todo/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SigninDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

func SignInController(ctx *gin.Context) {
	var body SigninDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	db, err := db.GetDB()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	var user models.User
	tx := db.Where(&models.User{Username: body.Username}).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(400, gin.H{
				"message": "Invalid Credential",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"message": tx.Error.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Credential",
		})
		return
	}

	// Generate JWT
	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Send JWT back to user
	ctx.JSON(201, gin.H{
		"token": token,
	})
}
