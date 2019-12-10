package main

import (
	"context"
	"flag"
	"fmt"
	"golang/series_3/src/09_the_context_package/log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8000", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100)) // to avoid this overwrite we changed requestIDKey = 42 to requestIDKey = key(42)

	log.Println(ctx, "Handler started...")
	defer log.Println(ctx,"Handler ended...")

	fmt.Printf("value for foo is %v\n", ctx.Value("foo"))
	//time.Sleep(5 * time.Second)
	//fmt.Fprintln(w, "hello")

	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <- ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
