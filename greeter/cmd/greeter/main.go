package main

import (
	"log"
	"net/http"
	"os"

	"greeter/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.GetGreeting)

	log.Printf("greeter listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
