package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	// mux.Handle("/public", http.StripPrefix("/public/", fs))
	mux.Handle("/", fs)
	log.Print("Listening on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
