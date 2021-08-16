package middlewares

import "github.com/thetkpark/golang-todo/services"

type Middleware struct {
	jwtManager *services.JWTManager
}

func NewMiddleware(jwtManager *services.JWTManager) *Middleware {
	return &Middleware{jwtManager: jwtManager}
}
