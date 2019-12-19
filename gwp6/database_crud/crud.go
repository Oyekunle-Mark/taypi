package main

import "database/sql"

type Post struct {
	ID      int
	Content string
	Author  string
}

var Db *sql.DB
