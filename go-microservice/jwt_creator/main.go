package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"client":     "Binh Le",
		"aud":        "billing.jwt.go.io",
		"iss":        "jwtgo.io",
		"exp":        time.Now().Add(time.Minute * 5).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	if err != nil {
		fmt.Println("Failed to generate JWT token")
		return
	}
	fmt.Println(validToken)
	fmt.Fprintf(w, string(validToken))
}

func handleRequests() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
