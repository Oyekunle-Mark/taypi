package main

import (
	"fmt"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola!")
}

func worldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/hello", helloFunc)
	http.HandleFunc("/world", worldFunc)

	server.ListenAndServe()
}
