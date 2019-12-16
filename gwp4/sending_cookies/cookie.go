package main

import "net/http"

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

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/cookie", sendCookie)
	server.ListenAndServe()
}
