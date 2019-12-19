package main

import "database/sql"

// Post is the post struct
type Post struct {
	ID      int
	Content string
	Author  string
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

func GetPost(id int) (post Post, err error) {
	post = Post{}

	query := `
		SELECT
		id, content, author
		FROM posts
		WHERE id = $1
	`
	err = Db.QueryRow(query, id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

func main() {

}
