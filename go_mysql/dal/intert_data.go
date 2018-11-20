package dal

import (
	"database/sql"
	"fmt"

	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func InsertRow(file_location string) {

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		fmt.Println("Failed to connect DataBase.", err)
		os.Exit(100)
	}

	uuid, err := exec.Command("uuidgen").Output()

	f, err := ioutil.ReadFile(strings.TrimSpace(file_location))

	if err != nil {
		fmt.Println("Failed to read the file")
		os.Exit(100)
	}

	rows, err := db.Exec(fmt.Sprintf("INSERT INTO testtable(uuid, random_data) VALUES ('%s', '%s')", string(uuid), f))
	_ = rows

	if err != nil {
		fmt.Println("Failed to insert data. ", err)
	} else {
		fmt.Println("Data Inserted.")
	}

}
