package dal

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

/*
psql -U postgres -h 172.16.205.131 -d testdb
*/

func Connect() {

	options := &pg.Options{
		User:     "postgres",
		Addr:     "172.16.205.131:5432",
		Database: "test",
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Println("Failed to connect database.")
		os.Exit(100)
	}

	log.Println("Connected ......")

	CreateTestTable(db)

	closeErr := db.Close()

	if closeErr != nil {
		log.Println("Failed to close DB connection.")
		os.Exit(100)
	}

	log.Println("Connection closed ....")
	return
}
