package main

import (
	"fmt"
	"time"
)

func callerA(ch chan string) {
	ch <- "Hello, World!"
}

func callerB(ch chan string) {
	ch <- "Hola, Mundo!"
}

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go callerA(ch1)
	go callerB(ch2)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)

		select {
		case msg := <-ch1:
			fmt.Printf("%s from A\n", msg)
		case msg := <-ch2:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("Default")
		}
	}
}
