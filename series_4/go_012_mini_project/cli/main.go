/*
Application - Reads data from storage, runs complex computation,
and writes results back to storage.
 */
package main

import (
	"log"
	"golang/series_4/go_012_mini_project/storage"
	"golang/series_4/go_012_mini_project/compute"
	// "golang/series_4/go_012_mini_project/silly" //imported and not used: "golang/series_4/go_012_mini_project/silly"
	_ "golang/series_4/go_012_mini_project/silly" // will just call init function of silly package, and will nots be used here in cli package
)

func init() {
}

func main() {

	//	TODO 1 - Get data from storage
	data, err := storage.GetData()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(data)

	// TODO 2 - Run complex computation
	result := compute.RunComputation(data)
	log.Println("Result of runComputation: ", result)

	// TODO 3 - Save result to storage
	if err = storage.StoreResult(result); err != nil {
		log.Fatalln(err)
	}
}

/*
series_4/go_012_mini_project$ go run main.go
2019/02/10 00:34:19 Package `storage` being initialized
2019/02/10 00:34:19 Package `compute` being initialized
2019/02/10 00:34:19 0xc00000e1f0
2019/02/10 00:34:19 Result of runComputation:  8
2019/02/10 00:34:19 Failed to get storage engine
exit status 1
series_4/go_012_mini_project$ go run main.go
2019/02/10 00:34:22 Package `storage` being initialized
2019/02/10 00:34:22 Package `compute` being initialized
2019/02/10 00:34:22 No data found
exit status 1
 */