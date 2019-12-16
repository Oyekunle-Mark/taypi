package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello World")

	cookie := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(message),
	}

	http.SetCookie(w, &cookie)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("flash")

	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		newCookie := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}

		http.SetCookie(w, &newCookie)

		decodedMsg, _ := base64.URLEncoding.DecodeString(cookie.Value)
		fmt.Fprintln(w, string(decodedMsg))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/set", setMessage)
	http.HandleFunc("/show", showMessage)

	server.ListenAndServe()
}
