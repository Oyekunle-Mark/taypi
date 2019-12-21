package main

import (
	"net/http"

	"github.com/Oyekunle-Mark/taypi/gwp7/rest_crud/request"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/post/", request.HandleRequest)
	server.ListenAndServe()
}
