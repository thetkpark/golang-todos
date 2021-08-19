package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thetkpark/golang-todo/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SignInController auth
// @Summary Login the user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body controllers.UserCredentialDto true "username and password to login"
// @Success 201 {object} controllers.TokenResponseDto "the user is login and token is given back"
// @Failure 400 {object} controllers.ErrorMessage "Missing some attribute or invalid credential"
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/signin [post]
func (c *Controller) SignInController(ctx *gin.Context) {
	var body UserCredentialDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, ErrorMessage{Message: err.Error()})
		return
	}

	var user models.User
	tx := c.db.Where(&models.User{Username: body.Username}).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(400, ErrorMessage{Message: "Invalid Credential"})
			return
		}
		ctx.JSON(500, ErrorMessage{Message: tx.Error.Error()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(400, ErrorMessage{Message: "Invalid Credential"})
		return
	}

	// Generate JWT
	token, err := c.jwtManager.GenerateJWT(user.ID)
	if err != nil {
		ctx.JSON(500, ErrorMessage{Message: err.Error()})
		return
	}

	// Send JWT back to user
	ctx.JSON(201, TokenResponseDto{token})
}
