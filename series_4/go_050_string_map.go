package main

import "strings"

func main() {
	s := "abc-efg"

	result := strings.Map(mapFunc, s)
}


func mapFunc(in rune) rune {
	switch in {

	}
	return in
}