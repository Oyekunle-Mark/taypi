package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type world struct{}

func (h *world) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World!")
}

func main() {
	hello := hello{}
	world := world{}

	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.Handle("/hello", &hello)
	http.Handle("world", &world)

	server.ListenAndServe()
}
