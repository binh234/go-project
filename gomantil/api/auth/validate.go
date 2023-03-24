package auth

import (
	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
)

func Validate(tokenA string) string {

	token, err := jwt.ParseWithClaims(tokenA, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "token not valid"
	}

	claims, ok := token.Claims.(*Claims)

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return "token expired"
	}

	if ok && token.Valid {

		return "valid"
	} else {
		return "token not valid"
	}

}
