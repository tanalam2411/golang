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

}

func fixHeader(s string) string {
	s = strings.ToLower(s)
	return strings.Title(s)
}


func getRow(in <-chan string) (ch <-chan *Person) {
	c := make(chan *Person)

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
			}
		}
	}()
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
