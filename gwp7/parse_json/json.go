package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func unmarshal(filename string) (Post, error) {
	var post Post
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return post, err
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return post, err
	}

	json.Unmarshal(jsonData, &post)
	return post, nil
}

func main() {
	jsonFile, err := os.Open("post.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Error reading json data:", err)
		return
	}

	var post Post

	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
