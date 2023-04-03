package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

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
