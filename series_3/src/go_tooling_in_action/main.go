// package  related documentation which you can see on `$ go list -f '{{ .Doc }}'`
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

// _ is for ignoring unsed package pprof

var re = regexp.MustCompile("^(.+)@golang.org$")

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serving at http://localhost:8000 ...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1) // -1 means match all

	if match != nil {
		fmt.Fprintf(w, "Hello, gopher %s\n", match[0][1])
		return
	}

	fmt.Fprintln(w, "Hello, world !!!")
}
