package main

import "fmt"

type Post struct {
	ID      int
	Content string
	Author  string
}

var PostByID map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostByID[post.ID] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostByID = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{ID: 1, Content: "First content.", Author: "Dante Cuavas"}
	post2 := Post{ID: 2, Content: "Second content.", Author: "Pirly Matti"}
	post3 := Post{ID: 3, Content: "Third content.", Author: "Dante Cuavas"}
	post4 := Post{ID: 4, Content: "Fourth content.", Author: "Rossi Manuel"}
	post5 := Post{ID: 5, Content: "Fifth content.", Author: "Dante Cuavas"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)
	store(post5)

	fmt.Println(PostByID[1])
	fmt.Println(PostByID[2])

	for _, post := range PostsByAuthor["Dante Cuavas"] {
		fmt.Println(post)
	}
}
