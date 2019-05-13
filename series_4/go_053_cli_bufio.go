package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type P struct {
	name   string
	age    int
	height float64
}

func main() {

	linesCh := readFile("data.csv")

	header := get_header(linesCh)

	fmt.Printf("%-15s %4s %8s\n",
		fixHeader(header[0]), fixHeader(header[1]),
		fixHeader(header[2]))

	personCh := get_row(linesCh)

	for p := range personCh {
		fmt.Printf("%-15v %4v %8.1f\n", p.name, p.age, p.height)
	}

}

func fixHeader(s string) string {
	s = strings.ToLower(s)
	return strings.Title(s)
}


func get_row(in <-chan string) (ch <-chan *P) {
	c := make(chan *P)

	go func() {
		var name string
		var age int
		var height float64

		for line := range in {

			fields := strings.Split(line, ",")

			switch len(fields) {

			case 3:
				fmt.Sscan(fields[2], &height)
				fallthrough
			case 2:
				fmt.Sscan(fields[1], &age)
				fallthrough
			case 1:
				name = strings.TrimSpace(fields[0])
				name = strings.Title(strings.ToLower(name))
			}

			c <- &P{name: name, age: age, height: height}
			name = ""
			age = 0
			height = 0.0
		}

		close(c)
	}()

	return c
}

func get_header(in <-chan string) (header []string) {
	input := <-in
	input = strings.TrimSpace(input)
	header = strings.Split(input, ",")
	return
}


func readFile(filename string) (ch <-chan string) {
	c := make(chan string)

	go func() {
		f, err := os.Open(filename)

		if err != nil {
			return
		}

		bio := bufio.NewReader(f)

		var input string

		for err == nil {
			input, err = bio.ReadString('\n')
			c <- input
		}

		close(c)

		f.Close()
	}()

	ch = c
	return
}
