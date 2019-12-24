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
