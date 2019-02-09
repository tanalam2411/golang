package compute

import "log"

func init() {
	log.Println("Package `compute` being initialized")
}

func RunComputation(d *string)int {
	return len(*d)
}
