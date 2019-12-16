package main

import (
	"encoding/base64"
	"net/http"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello World")

	cookie := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(message),
	}

	http.SetCookie(w, &cookie)
}

func main() {

}
