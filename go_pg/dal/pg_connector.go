package dal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectAndCreateTable() {

	connStr := "user=postgres host=172.16.205.131 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	fmt.Println(db, err)

	rows, err := db.Query(`CREATE TABLE TestTable (
		uuid  			varchar(100),
		timestamp   	timestamp DEFAULT current_timestamp,
		random_data 	bytea

);`)

	if err != nil {
		fmt.Println("Failed to create table testtable. Reason: ", err)
		//os.Exit(100)
	}

	fmt.Println(rows, err)

	db.Close()

}
