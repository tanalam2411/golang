package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background()
	// Cancel the request if it takes more than 1 second to complete.
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "foo", "bar")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
