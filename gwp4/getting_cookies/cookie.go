package main

import (
	"fmt"
	"net/http"
)

func sendCookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "Oyekunle",
		Value:    "Go Developer",
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "Mark",
		Value:    "Backend Engineer",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header["Cookie"]

	fmt.Fprintln(w, cookie)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/cookie", sendCookie)
	http.HandleFunc("/getCookie", getCookie)
	
	server.ListenAndServe()
}
