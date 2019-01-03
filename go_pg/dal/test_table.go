package dal

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

// type TestTable struct {
// 	tableName   struct{}  `sql:"test_table`
// 	uuid        string    `sql:"uuid"`
// 	timestamp   time.Time `sql:"timestamp"`
// 	random_data []byte    `sql:"random_data"`
// }

type TestTable struct {
	uuid      string    `sql:"uuid"`
	timestamp time.Time `sql:"timestamp"`
}

func CreateTestTable(db *pg.DB) error {
	// options := &orm.CreateTableOptions{
	// 	IfNotExists: true,
	// }

	// createErr := db.CreateTable(&TestTable{}, options)

	for _, model := range []interface{}{&TestTable{}} {
		createErr := db.CreateTable(model, &orm.CreateTableOptions{

			IfNotExists: true,
		})
		log.Println(createErr)
		// panicIf(err)
	}

	// createErr := db.CreateTable(&TestTable{}, &orm.CreateTableOptions{
	// 	Temp:        true,
	// 	IfNotExists: true,
	// })

	// if createErr != nil {
	// 	log.Println("Error while creating table TestTable, Reason: %v", createErr)
	// 	return createErr
	// }

	log.Println("Table TestTable created .......")
	return nil
}
