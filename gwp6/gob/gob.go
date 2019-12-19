package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	
	if err != nil {
		panic(err)
	}
}
