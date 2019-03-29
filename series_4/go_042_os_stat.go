package main

import (
	"fmt"
	"log"
	"os"
)

func main() {



	for _, a := range os.Args[1:] {
		t := getFileType(a)
		fmt.Printf("%v = %v\n", a, t)
	}

}

func getFileType(f string) (t string) {

	fi, err := os.Stat(f)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	if fi.IsDir() {
		return "Is dir"
	}

	return "is file"

	//switch mode := fi.Mode(); {
	//case mode.IsDir():
	//	fmt.Println("is Dir")
	//case mode.IsRegular():
	//	fmt.Println("is file")
	//default:
	//	fmt.Println("!!!")

}