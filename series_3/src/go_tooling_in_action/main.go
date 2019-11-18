// package  related documentation which you can see on `$ go list -f '{{ .Doc }}'`
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path
	match := re.FindAllStringSubmatch(path, -1) // -1 means match all

	if match != nil {
		fmt.Fprintln(w, "Hello, gopher %s", match[1])
	}

	fmt.Fprintln(w, "Hello, world !!!")
}
