package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID int
	Email string
	Role string
	PasswordHash string
}

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
