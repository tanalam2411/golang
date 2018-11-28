package main

import (
	dal "./dal"
	"bufio"
	"fmt"
	"os"
)

/*
https://stackoverflow.com/a/50390826
root@:~/performance/golang # go get -u github.com/golang/dep/cmd/dep
https://medium.com/@wembleyleach/how-to-use-the-official-mongodb-go-driver-9f8aff716fdb

*/

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter absolute path of file(utf-8 format): ")
	file_location, _ := reader.ReadString('\n')
	dal.InsertRow(file_location)

}
