package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Post is the structure of the XML file
type Post struct {
	XMLName  xml.Name  `xml:"post"`
	ID       int       `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Comments []Comment `xml:"comments>comment"`
	XML      string    `xml:",innerxml"`
}

// Author is the author element
type Author struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}

// Comment is the comment element
type Comment struct {
	ID      int    `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")

	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}

	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var post Post

	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
