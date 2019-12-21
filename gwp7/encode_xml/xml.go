package main

import (
	"encoding/xml"
	"fmt"
	"os"
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

	xmlFile, err := os.Create("post.xml")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)

	if err != nil {
		fmt.Println("Error encoding xml to file:", err)
		return
	}
}
