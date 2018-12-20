package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

/*

https://www.youtube.com/watch?v=CIh8qN7LO8M
$ go get github.com/gorilla/websocket

*/

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request))
}
