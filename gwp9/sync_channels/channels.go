package main

import (
	"fmt"
	"time"
)

func printNumbers(ch chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}

	ch <- true
}

func printLetters(ch chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}

	ch <- true
}

func main() {
	ch := make(chan bool)

	go printNumbers(ch)
	go printLetters(ch)
	
	<-ch
	<-ch
}
