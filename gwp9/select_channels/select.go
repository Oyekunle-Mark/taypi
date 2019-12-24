package main

import (
	"fmt"
	"time"
)

func callerA(ch chan string) {
	ch <- "Hello, World!"
	close(ch)
}

func callerB(ch chan string) {
	ch <- "Hola, Mundo!"
	close(ch)
}

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go callerA(ch1)
	go callerB(ch2)

	var msg string
	ok1, ok2 := true, true

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)

		select {
		case msg, ok1 = <-ch1:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-ch2:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}
