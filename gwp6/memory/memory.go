package main

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
	
}