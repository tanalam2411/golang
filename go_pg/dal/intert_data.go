package dal

import (
	"database/sql"
	"fmt"

	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	_ "github.com/lib/pq"
)

func InsertRow(file_location string) {

	connStr := "user=postgres host=172.16.205.131 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	uuid, err := exec.Command("uuidgen").Output()

	// log.Println(err, uuid)
	// fmt.Printf("-----------------------\n")
	// fmt.Println(string(uuid))

	// f, err := os.Open("/home/tan/Pictures/Wallpapers/machine_learning.png")
	// f, err := ioutil.ReadFile("/home/tan/tanveer/golang/go_pg/dal/input.txt")
	// f, err := ioutil.ReadFile("/home/tan/Downloads/output.csv")
	fmt.Println(file_location)
	f, err := ioutil.ReadFile(strings.TrimSpace(file_location))

	if err != nil {
		fmt.Println("Failed to read image file")
		os.Exit(100)
	}

	// rows, err := db.QueryRow(`INSERT INTO testtable(uuid, random_data) VALUES($1, $2)`, string(uuid), f)

	x := fmt.Sprintf("INSERT INTO testtable(uuid, random_data)  VALUES ('%s', '%s')", string(uuid), f)
	fmt.Println(x)

	rows, err := db.Exec(fmt.Sprintf("INSERT INTO testtable(uuid, random_data)  VALUES ('%s', '%s')", string(uuid), f))

	fmt.Println(rows, err)
}
