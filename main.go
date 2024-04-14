package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL)
		w.Write([]byte("Hello World"))
	})
	log.Print("Listing on port :8080")
	http.ListenAndServe(":8080", mux)
}
