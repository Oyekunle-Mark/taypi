package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	file, _, err := r.FormFile("uploaded")

	if err == nil {
		data, err := ioutil.ReadAll(file)

		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/parse", parse)
	server.ListenAndServe()
}
