package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome!")

	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
