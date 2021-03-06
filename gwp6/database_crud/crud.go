package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Post is the post struct
type Post struct {
	ID       int
	Content  string
	Author   string
	Comments []Comment
}

// Comment added to posts
type Comment struct {
	ID      int
	Content string
	Author  string
	Post    *Post
}

// Db instance
var Db *sql.DB

func init() {
	var err error

	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")

	if err != nil {
		panic(err)
	}
}

// Create method on Comment type creates a new comment
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
	}

	query := `
		INSERT INTO 
		comments (content, author, post_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err = Db.QueryRow(query, comment.Content, comment.Author, comment.Post.ID).Scan(&comment.ID)
	return
}

// Posts returns a number of post based on limit
func Posts(limit int) (posts []Post, err error) {
	query := `
		SELECT
		id, content, author
		FROM posts
		LIMIT $1
	`
	rows, err := Db.Query(query, limit)

	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Content, &post.Author)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	rows.Close()
	return
}

// GetPost gets all the posts in the db
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}

	query := `
		SELECT
		id, content, author
		FROM posts
		WHERE id = $1
	`
	err = Db.QueryRow(query, id).Scan(&post.ID, &post.Content, &post.Author)

	query = `
		SELECT
		id, content, author
		FROM comments
		WHERE post_id = $1
	`
	rows, err := Db.Query(query, id)

	for rows.Next() {
		comment := Comment{Post: &post}

		err = rows.Scan(&comment.ID, &comment.Content, &comment.Author)

		if err != nil {
			return
		}

		post.Comments = append(post.Comments, comment)
	}

	rows.Close()
	return
}

// Create adds a post to the db
func (post *Post) Create() (err error) {
	query := `
		INSERT
		INTO posts
		(content, author)
		VALUES ($1, $2)
		RETURNING id
	`
	stmt, err := Db.Prepare(query)

	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update updates the content of a post in the db
func (post *Post) Update() (err error) {
	query := `
		UPDATE posts
		SET content = $2, author = $3
		WHERE id = $1
	`
	_, err = Db.Exec(query, post.ID, post.Content, post.Author)
	return
}

// Delete removes a post from the db
func (post *Post) Delete() (err error) {
	query := `
		DELETE
		FROM posts
		WHERE id = $1
	`
	_, err = Db.Exec(query, post.ID)
	return
}

func main() {
	post := Post{Content: "...my soldiers remain eternal.", Author: "Tu Pac"}
	post.Create()
	fmt.Println(post)

	comment := Comment{Content: "Hot lines right there.", Author: "Jama Man", Post: &post}
	comment.Create()

	readPost, _ := GetPost(post.ID)
	fmt.Println(readPost)

	readPost.Content = "Get rich or die trying..."
	readPost.Author = "50 Cent"
	err := readPost.Update()

	if err != nil {
		panic(err)
	}

	posts, _ := Posts(1)
	fmt.Println(posts)

	readPost.Delete()
}
