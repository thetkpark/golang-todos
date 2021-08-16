package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

func (c *Controller) RegisterController(ctx *gin.Context) {
	var bodyData RegisterDto
	if err := ctx.ShouldBindJSON(&bodyData); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bodyData.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Username: bodyData.Username,
		Password: string(hashedPassword),
	}

	// Check exising username
	var existingUser int64
	tx := c.db.Model(&models.User{}).Where(&models.User{Username: bodyData.Username}).Count(&existingUser)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(existingUser)
	if existingUser != 0 {
		ctx.JSON(400, gin.H{
			"message": "Username is not available",
		})
		return
	}

	// Create new user
	if tx := c.db.Create(&user); tx.Error != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Generate JWT
	token, err := c.jwtManager.GenerateJWT(user.ID)
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
