package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			// Parse takes the token string and a function for looking up the key. The latter is especially
			// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
			// head of the token to identify which key to use, but the parsed token (head and claims) is provided
			// to the callback, providing flexibility.
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return MySigningKey, nil
			})

			if err != nil {
				fmt.Fprint(w, err.Error())
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if claims["aud"] != "billing.jwtgo.io" {
					fmt.Fprintf(w, "Invalid aud")
					return
				}
				if claims["iss"] != "jwtgo.io" {
					fmt.Fprintf(w, "Invalid iss")

				}
				if token.Valid {
					endpoint(w, r)
				}
			} else {
				fmt.Fprint(w, err.Error())
			}
		} else {
			fmt.Fprintf(w, "No authorization token provided")
		}
	})
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret information")

}

func handleRequests() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	fmt.Printf("server")
	handleRequests()
}
