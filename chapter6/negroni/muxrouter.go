package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome!")

	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
