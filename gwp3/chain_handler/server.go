package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("Hanlder function called: %s\n", name)

		h.ServeHTTP(w, r)
	})
}

func main() {
	hello := helloHandler{}

	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.Handle("/hello", log(hello))

	server.ListenAndServe()
}
