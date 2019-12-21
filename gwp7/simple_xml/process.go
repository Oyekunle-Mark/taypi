package main

import "encoding/xml"

type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      int      `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XML     string   `xml:",innerxml"`
}

type Author struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}
