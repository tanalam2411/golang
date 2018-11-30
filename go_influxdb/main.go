package main

import (
	dal "./dal"
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter absolute path of file(utf-8 format): ")
	file_location, _ := reader.ReadString('\n')
	dal.InsertRow(file_location)

}
