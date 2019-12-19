package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	allPosts := []Post{
		Post{1, "Mickey rocks", "Walt Disnery"},
		Post{2, "Tom is after Jerry", "Warner Bros."},
		Post{3, "Who run the world?", "Sharks"},
	}

	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)

		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}
