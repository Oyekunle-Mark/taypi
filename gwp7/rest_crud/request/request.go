package request

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/Oyekunle-Mark/taypi/gwp7/rest_crud/data"
)

// HandleRequest maps the request to the right function
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	post, err := data.Retrieve(id)

	if err != nil {
		return
	}

	jsonData, err := json.MarshalIndent(&post, "", "\t")

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var post data.Post
	json.Unmarshal(body, &post)

	err = post.Create()

	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	post, err := data.Retrieve(id)

	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	json.Unmarshal(body, &post)
	err = post.Update()

	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	post, err := data.Retrieve(id)

	if err != nil {
		return
	}

	err = post.Delete()

	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}
