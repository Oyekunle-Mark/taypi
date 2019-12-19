package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")

	err := ioutil.WriteFile("data1", data, 0644)

	if err != nil {
		panic(err)
	}

	readData, _ := ioutil.ReadFile("data1")
	fmt.Println("Read file data1:\n", string(readData))

	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file1\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	readData2 := make([]byte, len(data))
	bytes, _ = file2.Read(readData2)

	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(readData2))
}
