package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (t *Token) SetToken(ctx context.Context) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	claims := Claims{
		"binh",
		"binhnd234@gmail.com",
		"788949993",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "Binh Le",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	return signedToken
}
