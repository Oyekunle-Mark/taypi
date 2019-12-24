package main

import (
	"fmt"
	"time"
)

func thrower(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Threw >> %d\n", i)
	}
}

func catcher(ch chan int) {
	for i := 0; i < 5; i++ {
		msg := <-ch
		fmt.Printf("Caught << %d\n", msg)
	}
}

func main() {
	ch := make(chan int)

	go thrower(ch)
	go catcher(ch)

	time.Sleep(100 * time.Millisecond)
}
