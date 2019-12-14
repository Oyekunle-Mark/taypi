package main

import (
	"fmt"
	"net/http"
)

type resHandler struct{}

func (h *resHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	handler := resHandler{}

	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.Handle("/hello", &handler)

	server.ListenAndServe()
}
