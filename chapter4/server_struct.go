package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang is Awesome!")
}

func main() {
	http.HandleFunc("/go", messageHandler)

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening ...")
	server.ListenAndServe()
}
