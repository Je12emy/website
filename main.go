package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	// mux.Handle("/public", http.StripPrefix("/public/", fs))
	mux.Handle("/", fs)
	log.Printf("Listening on port :%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
