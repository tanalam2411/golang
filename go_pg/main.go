package main

import (
	dal "./dal"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Select options: ")
	fmt.Println("To create table(testtable) enter : 1 ")
	fmt.Println("To insert data in table testtable: 2: ")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your option(1 or 2) : ")
	input_option, _ := reader.ReadString('\n')

	fmt.Println(input_option)

	if strings.Compare(strings.TrimSpace(input_option), "1") == 0 {
		fmt.Println("Executing option 1.")
		dal.ConnectAndCreateTable()
	} else {
		fmt.Println("Enter absolute path of file(utf-8 format): ")

		file_location, _ := reader.ReadString('\n')

		dal.InsertRow(file_location)
	}

	// dal.ConnectAndCreateTable()

	// dal.InsertRow()

}
