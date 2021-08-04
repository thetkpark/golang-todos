package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/thetkpark/golang-todo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := authHeader[len(BearerSchema):]
		token, err := services.ValidateJWT(tokenString)
		if token != nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims["user_id"], claims)
			c.Set("userId", claims["user_id"])
			c.Next()
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
