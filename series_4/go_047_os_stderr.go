package main

import "os"

func main() {
	_, _ = os.Stdout.WriteString("Writing to os.Stdout \n")
	_, _ = os.Stderr.WriteString("Writing to os.Stderr \n")

}