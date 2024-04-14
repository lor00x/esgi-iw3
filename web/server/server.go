package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Nouveau serveur
	s := http.NewServeMux()

	// Nouvelle route
	s.HandleFunc("/hello", helloWorld)

	// Start !
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
