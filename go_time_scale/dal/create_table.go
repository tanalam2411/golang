package dal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectAndCreateTable() {

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Failed to connect DataBase.", err)
		os.Exit(100)
	}

	rows, err := db.Query(`CREATE TABLE TestTable (
		uuid  			varchar(100),
		timestamp   	timestamp DEFAULT current_timestamp,
		random_data 	bytea

);`)

	if err != nil {
		fmt.Println("Failed to create table testtable. Reason: ", err)
		os.Exit(100)
	}

	_ = rows

	db.Close()

}
