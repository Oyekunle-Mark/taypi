package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	handler := hello{}

	server := http.Server{
		Addr:    "127.0.0.1:5000",
		Handler: &handler,
	}

	server.ListenAndServe()
}
