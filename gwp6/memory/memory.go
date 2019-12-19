package main

type Post struct {
	ID      int
	Content string
	Author  string
}

var PostByID map[int]*Post
var PostByAuthor map[string][]*Post
