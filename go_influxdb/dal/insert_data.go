/*
https://stackoverflow.com/questions/38067435/how-to-write-to-continuosly-write-to-influxdb-using-golang-client
https://github.com/influxdata/influxdb/blob/master/client/example_test.go
https://github.com/influxdata/influxdb/issues/756

go get github.com/influxdata/influxdb/client/v2

*/

package dal

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

func InsertRow(file_location string) {

	httpClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     connStr,
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer httpClient.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	in_uuid, err := exec.Command("uuidgen").Output()
	f, err := ioutil.ReadFile(strings.TrimSpace(file_location))

	tags := map[string]string{}

	fields := map[string]interface{}{
		"Uuid":      fmt.Sprintf("%s", in_uuid),
		"Timestamp": time.Now(),
		"Blob_data": fmt.Sprintf("%s", f),
	}

	pt, err := client.NewPoint("test_collection", tags, fields, time.Now())

	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := httpClient.Write(bp); err != nil {
		log.Fatal(err)
	}

	if err := httpClient.Close(); err != nil {
		log.Fatal(err)
	}

}
