package main

import (
	"fmt"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	handler := handler{}

	server := http.Server{
		Addr:    "127.0.0.1:5000",
		Handler: &handler,
	}

	server.ListenAndServe()
}
