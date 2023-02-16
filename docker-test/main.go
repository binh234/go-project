package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello User")
}

func getHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi page")
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/hi", getHi)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Error: Could not start server at port 8081 %v\n", err)
	}
}
