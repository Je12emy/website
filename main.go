package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	// mux.Handle("/public", http.StripPrefix("/public/", fs))
	mux.Handle("/", fs)
	log.Print("Listening on port :8080")
	http.ListenAndServe(":8080", mux)
}
