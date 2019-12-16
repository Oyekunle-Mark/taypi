package main

import (
	"fmt"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	fmt.Fprintln(w, "(3)", r.PostForm)
	fmt.Fprintln(w, "(4)", r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/parse", parse)
	server.ListenAndServe()
}
