package main

import (
	"fmt"
	"log"
	"net/http"
)

//Custom Handler
type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

//Ordinary function as handler
func messageHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang looking goood.")
}

//Handler logic into a Closure
func messageHandler3(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("chapter4/staticweb/public"))
	mux.Handle("/", fs)

	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)

	mh2 := &messageHandler{"net/http is awesome"}
	mux.Handle("/message", mh2)

	//Conver the messageHandler function to a HandlerFunc type
	mh3 := http.HandlerFunc(messageHandler2)
	mux.Handle("/golang", mh3)

	mux.Handle("/python", messageHandler3("Python is Awesome too."))
	mux.Handle("/learning", messageHandler3("Learning new things is a must"))

	//Use the shortcut method ServeMus.HandleFunc
	mux.HandleFunc("/golang2", messageHandler2)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
