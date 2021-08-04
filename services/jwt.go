package services

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type JWTPayload struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		log.Fatalln("JWT_SECRET env is required")
	}
	return secret
}

func GenerateJWT(userId uint) (string, error) {
	claims := &JWTPayload{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(getSecret()))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return tokenString, nil
}

func ValidateJWT(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(getSecret()), nil
	})
}
