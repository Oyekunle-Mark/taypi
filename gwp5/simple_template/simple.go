package main

import (
	"net/http"
	"text/template"
)

func simple(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("simple.html")
	t.Execute(w, "Very basic template!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/simple", simple)
	server.ListenAndServe()
}
