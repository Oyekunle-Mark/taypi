package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("Hanlder function called: %s\n", name)

		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}

	http.HandleFunc("/hello", log(helloFunc))

	server.ListenAndServe()
}
