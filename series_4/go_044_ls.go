//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//var longListing = false
//
//
//func main() {
//	offset := 1
//
//	if len(os.Args) >= 2 && os.Args[offset] == "-l" {
//		longListing = true
//		offset++
//	}
//
//	files := os.Args[offset:]
//
//	if len(files) == 0 {
//		files = []string{"."}
//	}
//
//	if longListing {
//		showLongListing(files)
//		return
//	}
//
//	showShortListing(files)
//}
//
//
//func showShortListing(files []string) {
//	var noFilesList []string
//	var filesList []string
//	var dirListing []string
//
//	for _, f := range files {
//		fi, err := os.Stat(f)
//
//		if err != nil {
//			s := fmt.Sprintf("ls: %v: no file or direvtory", f)
//			noFilesList = append(noFilesList, s)
//			continue
//		}
//	}
//}





package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Println("Hello, playground") // this gets captured

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Printf("Captured: %s", out) // prints: Captured: Hello, playground
}