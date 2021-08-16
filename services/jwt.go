package services

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JWTManager struct {
	secret   string
	duration time.Duration
}

func NewJWTManager(secret string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secret:   secret,
		duration: duration,
	}
}

type JWTPayload struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func (manager *JWTManager) GenerateJWT(userId uint) (string, error) {
	claims := &JWTPayload{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(manager.secret))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return tokenString, nil
}

func (manager *JWTManager) ValidateJWT(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(manager.secret), nil
	})
}
