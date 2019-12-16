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

	w.Header().Add("Set-Cookie", cookie1.String())
	w.Header().Set("Set-Cookie", cookie2.String())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/cookie", sendCookie)
	server.ListenAndServe()
}
