package main

import (
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("tmpl.html")
	// t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}