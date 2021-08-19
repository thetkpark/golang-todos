package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserCredentialDto auth
type UserCredentialDto struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

// RegisterController auth
// @Summary Register the new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body controllers.UserCredentialDto true "username and password to create user"
// @Success 201 {object} controllers.TokenResponseDto "the user is register and token is given back"
// @Failure 400 {object} controllers.ErrorMessage "Missing some attribute or username is in used"
// @Failure 500 {object} controllers.ErrorMessage "Internal Server Error"
// @Router /api/regis [post]
func (c *AuthController) RegisterController(ctx *gin.Context) {
	var bodyData UserCredentialDto
	if err := ctx.ShouldBindJSON(&bodyData); err != nil {
		ctx.JSON(400, ErrorMessage{Message: err.Error()})
		return
	}

	// Check exising username
	existingUser, err := c.userRepository.FindByUsername(bodyData.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No user found, do nothing
		} else {
			ctx.JSON(500, ErrorMessage{Message: err.Error()})
			return
		}
	} else if existingUser != nil {
		ctx.JSON(400, ErrorMessage{Message: fmt.Sprintf("username %s is unavaliable", bodyData.Username)})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bodyData.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, ErrorMessage{Message: err.Error()})
		return
	}

	// Create new user
	user, err := c.userRepository.Create(bodyData.Username, string(hashedPassword))
	if err != nil {
		ctx.JSON(500, ErrorMessage{Message: err.Error()})
		return
	}

	// Generate JWT
	token, err := c.jwtManager.GenerateJWT(user.ID)
	if err != nil {
		ctx.JSON(500, ErrorMessage{Message: err.Error()})
		return
	}

	// Send JWT back to user
	ctx.JSON(201, TokenResponseDto{Token: token})
}

// TokenResponseDto auth
type TokenResponseDto struct {
	Token string `json:"token"`
}
