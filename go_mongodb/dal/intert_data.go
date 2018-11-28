package dal

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Data struct {
	Uuid      []byte
	Timestamp time.Time
	Blob_data []byte
}

func InsertRow(file_location string) {

	client, err := mongo.Connect(context.Background(), connStr, nil)
	if err != nil {
		fmt.Println("Failed to connect mongod server.", err)
		os.Exit(100)
	}

	db := client.Database("testdb")
	test_collection := db.Collection("test_collection")

	in_uuid, err := exec.Command("uuidgen").Output()
	f, err := ioutil.ReadFile(strings.TrimSpace(file_location))

	itemWrite := Data{Uuid: in_uuid, Timestamp: time.Now(), Blob_data: f}

	// fmt.Printf("itemWrite = %v\n", itemWrite)

	result, err := test_collection.InsertOne(context.Background(), itemWrite)
	if err != nil {
		fmt.Println("Failed to insert..", err)
		os.Exit(100)
	}
	fmt.Printf("result = %#v\n", result)

}
