package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func decode(filename string) (Post, error) {
	var post Post
	jsonFile, err := os.Open("post.json")

	if err != nil {
		fmt.Println("Error opening the JSON file:", err)
		return post, err
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)

	if err != nil {
		fmt.Println("Error decoding JSON file:", err)
		return post, err
	}

	return post, nil
}

func main() {
	post, err := decode("post.json")

	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	fmt.Println(post)
}
