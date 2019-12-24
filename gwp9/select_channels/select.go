package main

func callerA(ch chan string) {
	ch <- "Hello, World."
}


