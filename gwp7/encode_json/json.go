package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Post is the structure of the json file
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Author is the author object
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment is the comment object
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		ID:      1,
		Content: "Hello World!",
		Author: Author{
			ID:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				ID:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				ID:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	jsonData, err := json.MarshalIndent(&post, "", "\t")

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("post.json", jsonData, 0644)

	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
