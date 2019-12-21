package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Post is the structure of the XML file
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

// Author is the author element
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: " Hello World!", Author: Author{
			ID:   "2",
			Name: "Sau Sheong",
		},
	}

	xmlData, err := xml.MarshalIndent(&post, "", "\t")

	if err != nil {
		fmt.Println("Error marshalling to xml:", err)
		return
	}

	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(xmlData)), 0644)

	if err != nil {
		fmt.Println("Error writing xml to file:", err)
		return
	}
}
