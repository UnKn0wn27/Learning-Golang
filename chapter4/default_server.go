package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang looking goood.")
}

func main() {
	http.HandleFunc("/go", messageHandler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
