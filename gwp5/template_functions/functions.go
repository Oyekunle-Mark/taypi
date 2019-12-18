package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2011-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	tmpl := template.New("tmpl.html").Funcs(funcMap)
	t, _ := tmpl.ParseFiles("tmpl.html")

	t.Execute(w, time.Now())
}
