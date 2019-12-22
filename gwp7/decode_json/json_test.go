package main

import "testing"

func TestDecode(t *testing.T) {
	post, err := decode("post.json")

	if err != nil {
		t.Error(err)
	}

	id := 1

	if post.ID != id {
		t.Errorf("Wrong ID. Got %d, was expecting %d", post.ID, id)
	}

	content := "Hello World!"

	if post.Content != content {
		t.Errorf("Wrong post content. Got %s, was expecting %s", post.Content, content)
	}
}
