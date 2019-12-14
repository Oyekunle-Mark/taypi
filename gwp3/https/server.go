package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type resHandler struct{}

func (h *resHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	handler := resHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:5000",
		Handler: &handler,
	}

	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
