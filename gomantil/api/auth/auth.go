package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Key int

const MyKey Key = 0

// JWT schema of the data it will store
type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	jwt.StandardClaims
}

type Token struct {
	token string
}

func New() *Token {
	return &Token{}
}
