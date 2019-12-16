package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User   string
	Thread []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>Go Web Programming</title>
		</head>
		<body>
			<h1>Hello World!</h1>
		</body>
		</html>
	`

	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)

	fmt.Fprintln(w, "Service not available.")
}

func redirectExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http:www.google.com")
	w.WriteHeader(302)
}

func writeJSONExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post := Post{
		User:   "Bruce Lee",
		Thread: []string{"Fight", "Dance", "Sing"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", redirectExample)
	http.HandleFunc("/json", writeJSONExample)

	server.ListenAndServe()
}
