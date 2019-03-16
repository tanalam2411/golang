package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Questions = []string{
	"question1",
	"question2",
	"question3",
	"question4",
	"question5",
}

var Answers = map[string]string {
	"question1": "answer1",
	"question2": "answer2",
	"question3": "answer3",
	"question4": "answer4",
	"question5": "answer5",
}


type Request struct {
	q string
	c chan<- Response
}

type Response struct {
	q string
	a string
	d time.Duration
	serverId int
}

type ClientResponse struct {
	q string
	a string
}


const TotalServers = 5
const MAX_SERVER_TIME = 100


func init() {
	rand.Seed(time.Now().Unix())
}


func main() {
	qCh := client(Questions)
	ansCh := hub(qCh)

	for a := range ansCh {
		//fmt.Println(a)
		fmt.Printf("Question: %v\nAnswer  : %v\n", a.q, a.a)
	}
}


func hub(qCh <-chan string) <-chan ClientResponse {
	var wg sync.WaitGroup

	out := make(chan ClientResponse)

	go func() {
		for q := range qCh {
			ansCh := make(chan Response)
			req := Request{q: q, c:ansCh}

			for i := 0; i < TotalServers; i++ {
				go server(i, req)
			}

			wg.Add(1)
			go handler(out, q, ansCh, &wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}


func handler(out chan<- ClientResponse, theQuestion string, myAnsCh <-chan Response, wg *sync.WaitGroup) {
	timeout := time.After(MAX_SERVER_TIME / 2 * time.Nanosecond)

	cr := ClientResponse{q: theQuestion}

	select {
	case ans := <-myAnsCh:
		cr.a = ans.a
		out <- cr

	case <-timeout:
		cr.a = "Request timed out"
		out <- cr
	}

	wg.Done()
}


func server(id int, req Request) {
	d := time.Duration(rand.Intn(MAX_SERVER_TIME))
	time.Sleep(d * time.Nanosecond)

	a := Answers[req.q]
	resp := Response{q: req.q, a: a, d: d, serverId: id}
	req.c <- resp
}


func client(data []string) <-chan string {
	out := make(chan string)

	go func() {
		for _, q := range data {
			out <- q
		}
		close(out)
	}()

	return out
}