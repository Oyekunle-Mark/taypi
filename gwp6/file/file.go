package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("Hello World!\n")

	err := ioutil.WriteFile("data1", data, 0644)

	if err != nil {
		panic(err)
	}

	readData, _ := ioutil.ReadFile("data1")
	fmt.Println("Read file data1:\n", string(readData))
}
