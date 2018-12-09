package main

import (
	dal "./dal"
	"bufio"
	"fmt"
	"os"
)

/*
https://community.grafana.com/t/query-for-influxdb-to-select-a-range-of-tag-keys/10316/2
*/

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter absolute path of file(utf-8 format): ")
	file_location, _ := reader.ReadString('\n')
	dal.InsertRow(file_location)

}
