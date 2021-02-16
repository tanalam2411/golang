package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	compareBytes([]byte("abc"), []byte("bc"))
	compareStrings("abc", "bc")
	containsAny([]byte("abc"), []byte("cb"))
	hasPrefixByte([]byte("Golang"), []byte("Go"))
	joinBytes([]byte("Golang"), []byte("Go"))
}

func compareBytes(b1, b2 []byte) {

	res := bytes.Contains(b1, b2)
	fmt.Printf("compareBytes: %v - %v: %v\n", b1, b2, res)

	var b3 []byte

	b3 = append(b3, []byte("abcdef")...)
	fmt.Printf("b3: %v\n", b3)
}

func compareStrings(s1, s2 string) {
	res := strings.Contains(s1, s2)
	fmt.Printf("comapreBytes: %v - %v: %v\n", s1, s2, res)
}

func containsAny(b1, b2 []byte) {
	res := bytes.Contains(b1, b2)
	fmt.Printf("containsBytes: %v, %v: %v\n", b1, b2, res)
}

func hasPrefixByte(b1, b2 []byte) {
	res := bytes.HasPrefix(b1, b2)
	fmt.Printf("hasPrefixByte: %v, %v: %v\n", b1, b2, res)
}

func joinBytes(b1 []byte, b2 []byte) {
	var resB [][]byte
	resB = append(resB, b1)
	resB = append(resB, b2)
	res := bytes.Join(resB, []byte(" = "))
	fmt.Printf("joinBytes: %v, %v, %v: %v, %s\n", b1, b2, resB, res, res)

}
