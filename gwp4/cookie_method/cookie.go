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
	cookie, err := r.Cookie("Oyekunle")

	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie.")
	}

	cookies := r.Cookies()

	fmt.Fprintln(w, cookie)
	fmt.Fprintln(w, cookies)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/cookie", sendCookie)
	http.HandleFunc("/getCookie", getCookie)

	server.ListenAndServe()
}
