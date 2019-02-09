package storage

import (
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"time"

)

func init() {
	rand.Seed(time.Now().Unix())

	log.Println("Package `storage` being initialized")
}


// GetData returns if it available else returns error
func GetData() (*string, error) {

	i := rand.Intn(3)

	var d string

	switch i {
	case 0:
		d = "data one"
	case 1:
		d = "data two"
	default:
		return nil, errors.New("No data found")
	}
	return &d, nil
}


// StoreResult takes int and writes it to the storage engine
func StoreResult(r int) error {
	db, err := getStorageEngine()
	if err != nil {
		return err
	}

	*db = r
	return nil
}

func getStorageEngine() (*int, error) {
	i := rand.Intn(2)
	var d = 0
	if i==0 {
		return &d, nil
	}

	return nil, errors.New("Failed to get storage engine")
}