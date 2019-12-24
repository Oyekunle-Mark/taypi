package main

import "fmt"

func thrower(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Threw >> %d", i)
	}
}
