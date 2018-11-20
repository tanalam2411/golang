package dal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectAndCreateTable() {

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		fmt.Println("Failed to connect DataBase.", err)
		os.Exit(100)
	}

	rows, err := db.Query(`CREATE TABLE TestTable (
		uuid  			varchar(100),
		timestamp		DATETIME DEFAULT CURRENT_TIMESTAMP,
		random_data		BLOB

	);`)

	if err != nil {
		fmt.Println("Failed to create table testtable. Reason: ", err)
		os.Exit(100)
	} else {
		fmt.Println("testtable got created in database testdb")
	}

	_ = rows

	db.Close()

}
