package main

import (
	"fmt"
	"os"
)

var longListing = false


func main() {
	offset := 1

	if len(os.Args) >= 2 && os.Args[offset] == "-l" {
		longListing = true
		offset++
	}

	files := os.Args[offset:]

	if len(files) == 0 {
		files = []string{"."}
	}

	if longListing {
		showLongListing(files)
		return
	}

	showShortListing(files)
}


func showShortListing(files []string) {
	var noFilesList []string
	var filesList []string
	var dirListing []string

	for _, f := range files {
		fi, err := os.Stat(f)

		if err != nil {
			s := fmt.Sprintf("ls: %v: no file or direvtory", f)
			noFilesList = append(noFilesList, s)
			continue
		}

		if !fi.IsDir() {
			filesList = append(filesList, f)
			continue
		}

		dirListing = addDirListing(dirListing, f)
	}


	for _, s := range noFilesList {
		fmt.Println(s)
	}

	for _, s := range filesList {
		fmt.Println(s)
	}

	for _, s := range dirListing {
		fmt.Println(s)
	}

}


func addDirListing(listing []string, f string) []string {
	dir, err := os.Open(f)

	if err != nil {
		return listing
	}

	fileNames, err := dir.Readdirnames(0)

	if err != nil {
		return listing
	}

	listing = append(listing, "\n"+f+":")

	for _, d := range fileNames {
		listing = append(listing, d)
	}

	return listing
}


func showLongListing(files []string) {
	fmt.Println("showing long listing for: ", files)
}


func calSize(i int64) string {
	s := float64(i)
	unit := "B" //bytes

	if (s/1024) > 1.0 {
		s = s / 1024
		unit = "K"
	}

	if (s/1024) > 1.0 {
		s = s / 1024
		unit = "M"
	}

	if (s/1024) > 1.0 {
		s = s / 1024
		unit = "G"
	}

	if (s/1024) > 1.0 {
		s = s / 1024
		unit = "T"
	}

	return fmt.Sprintf("%5.2v%v", s, unit)
}




func permString(m os.FileMode) string {
	p := "rw-r--r--"
	return p
}
