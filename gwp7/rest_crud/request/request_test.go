package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Oyekunle-Mark/taypi/gwp7/rest_crud/data"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", HandleRequest)

	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v, expected 200.", writer.Code)
	}

	var post data.Post
	json.Unmarshal(writer.Body.Bytes(), &post)

	if post.ID != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	jsonData := strings.NewReader(`{
		"author": "Chief Oye",
		"content": "The end is upon us."
	}`)

	request, _ := http.NewRequest("PUT", "/post/1", jsonData)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v, expected 200.", writer.Code)
	}
}
