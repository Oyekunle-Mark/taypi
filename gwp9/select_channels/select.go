package main

func callerA(ch chan string) {
	ch <- "Hello, World."
}

func callerB(ch chan string) {
	ch <- "Hola, Mundo!"
}
