package main

import (
	"fmt"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/parse", parse)
	server.ListenAndServe()
}
