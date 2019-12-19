package main

import (
	"encoding/csv"
	"fmt"
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

	file, err := os.Open("posts.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	var posts []Post

	for _, item := range records {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{int(id), item[1], item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].ID)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
	fmt.Println(posts)
}
