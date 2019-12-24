package main

import (
	"fmt"
	"sync"
)

func thrower(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Threw >> %d\n", i)
	}

	wg.Done()
}

func catcher(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		msg := <-ch
		fmt.Printf("Caught << %d\n", msg)
	}

	wg.Done()
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go thrower(ch, &wg)
	go catcher(ch, &wg)

	wg.Wait()
}
