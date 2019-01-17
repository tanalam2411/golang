package main

import (
	"fmt"
)

/* AGENDA

- Maps
  - What are they?
  - Creating
  - Manipulation

*/

/*
Slices, maps and other functions can't be compared for equality check. So can't be used as key with in map
Maps are unordered
Are of reference type, meaning if assigned to other variable then a reference is created not a seperate copy
*/

func main() {

	deviceInventory := map[string]int{
		"vm":     10,
		"router": 4,
		"switch": 3,
	}

	fmt.Printf("%v, %T \n", deviceInventory, deviceInventory)
	fmt.Println(deviceInventory["vm"])

	// another way to declare map using make function

	deviceIPMap := make(map[string]string)

	deviceIPMap = map[string]string{"10.0.0.10": "dev_10"}

	//Add new key: val pair
	deviceIPMap["10.0.0.11"] = "dev_11"
	fmt.Printf("%v, %T \n", deviceIPMap, deviceIPMap)

	//Delete new key: val pair
	// delete: is not variadic function, os you can't delete more than 1 key:pair in a single func call
	delete(deviceIPMap, "10.0.0.11")
	fmt.Printf("%v, %T \n", deviceIPMap, deviceIPMap)

	// if key is not present in the map, it dosn't thorugh any error
	// to check key status use ok return value

	val1, ok1 := deviceIPMap["10.0.0.11"]
	// So if key is not present and value is of type string it will return empty string and ok as false
	// If value is of type int or float it will return 0 and false
	// using ok is an convention not mandetory
	fmt.Printf("%v, %T, %v \n", val1, val1, ok1) // , string, false

	val2, ok2 := deviceIPMap["10.0.0.10"]
	fmt.Printf("%v, %T, %v \n", val2, val2, ok2) // dev_10, string, true

	//Get lenght of map
	fmt.Println(len(deviceIPMap))

	// Assining to another variable creates a reference not a copy
	dupDeviceIPMap := deviceIPMap

	fmt.Println("deviceIPMap: ", deviceIPMap)
	fmt.Println("dupDeviceIPMap: ", dupDeviceIPMap)

	delete(dupDeviceIPMap, "10.0.0.10")

	fmt.Println("deviceIPMap: ", deviceIPMap)
	fmt.Println("dupDeviceIPMap: ", dupDeviceIPMap)

}

/* SUMMARY

- Maps
  - Collections of value types that are accessed via keys
  - Created via literals or via make function
  - Members accessed via [key] syntax
	- myMap["key"] = "value"
  - Check for presence with "value, ok" format of result -> actual value, bool
  - Reference type - Multiple assignments refer to same underlying data

*/
