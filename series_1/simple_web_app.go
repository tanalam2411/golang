package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func index_handler(w http.ResponseWriter, r *http.Request) {

	/*
		w - is http writter
		r *http.Request - reading through http request
	*/

	fmt.Fprintf(w, "Welcome to  Home page")

}

func about_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "about page")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Today's News", News: "This happened"}
	t, _ := template.ParseFiles("agg.html")
	t.Execute(w, p)
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil) // nil is like None
}
