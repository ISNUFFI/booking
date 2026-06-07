package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

const userIDKey = "user"

type User struct {
	ID           int	`json:"id"`
	Email        string	`json:"email"`
	Role         string	`json:"role"`
	passwordHash string
}

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
