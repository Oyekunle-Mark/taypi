package data

import (
	"database/sql"
	// the database driver
	_ "github.com/lib/pq"
)

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

// Retrieve returns a post of id id
func Retrieve(id int) (post Post, err error) {
	post = Post{}

	query := `
		SELECT *
		FROM posts
		WHERE id = $1
	`

	err = Db.QueryRow(query, id).Scan(&post.ID, &post.Content, &post.Author)
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
