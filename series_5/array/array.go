package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("--------- Arrays Introduction ------------")

	// Declaring an array
	var i[10]int
	fmt.Printf("i value: %v, type: %T\n", i, i)
	fmt.Printf("i Type: ", reflect.ValueOf(i).Kind())

	var f[10]float32
	fmt.Printf("f value: %v, type: %T\n", f, f)

	var s[10]string
	fmt.Printf("s value: %v, type: %T\n", s, s)

	var b[10]byte
	fmt.Printf("b value: %v, type: %T\n", b, b)

	var r[10]rune
	fmt.Printf("r value: %v, type: %T\n", r, r)


	// Initializing an array using an array literal
	var ii = [10]int{1, 2, 3}
	fmt.Printf("ii value: %v, type: %T\n", ii, ii)


	// short hand declaration
	iii := [5]int{4, 5, 6}
	fmt.Printf("iii value: %v, type: %T\n", iii, iii)


	// Letting go compiler infer the length of the array
	iiii := [...]int{7, 8, 9, 10}
	fmt.Printf("iiii value: %v, type: %T, len: %v\n", iiii, iiii, len(iiii))


	// Arrays are value type, so assigning an array variable to another will make copy of it, not reference.
	jjjj := iiii
	iiii[0] = 100
	fmt.Printf("value of iiii: %v and jjjj: %v\n", iiii, jjjj)

	// Passing reference of array
	kkkk := &jjjj
	jjjj[0] = 1000
	fmt.Printf("value of jjjj: %v, vaue of kkkk: %v\n", jjjj, kkkk)

	//
	iiiii := [...]int{2: 10, 5: 20}
	fmt.Println("Value of iiiii: ", iiiii)

	fmt.Println("---------- Iterating over array -----------")

	// Iterating over an array
	for i:=0; i < len(iiii); i++ {
		fmt.Printf("iiii: %v\n", iiii[i])
	}

	for i, v := range iiii {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	ind := 0
	for range iiii {
		fmt.Println("index + 1 iiii: ", iiii[ind])
		ind = ind + 1
	}

	fmt.Println("------------ multidimensional array ---------------")
	// multidimensional array
	ma := [2][2]int{
		{1, 2},
	}

	fmt.Printf("ma Type: %T, value: %v\n", ma, ma)

	ma1 := [...][4]int{{1, 2, 3}, {1, 2}, {1, 2, 3, 4}}
	fmt.Printf("ma1 Type: %T, ma1 value: %v\n", ma1, ma1)

	for _, par := range ma1{
		for _, child := range par {
			fmt.Println(par, child)
		}
	}

	fmt.Println("-------------- Array Comparision ---------------")
	a1 := [...]int{1, 2, 3}
	a2 := [...]int{1, 2 ,3}
	a3 := [...]int{1, 2, 4}

	fmt.Println("a1 == a2", a1==a2)
	fmt.Println("a1 == a3", a1==a3)

	fmt.Println("-------------- Check if item exists in array -----------------")
	ret := itemExists(a1, 1)
	fmt.Println("Item found: ", ret)

	 fmt.Println("------------- ")

	var ar1 [4]int
	var ar2 [4]int
	fmt.Println("is ar1 equal to ar2: ", reflect.DeepEqual(ar1, ar2))

	ar1[0] = 1
	fmt.Println("is ar1 equal to ar2: ", reflect.DeepEqual(ar1, ar2))

	var ar3_1 [4]int
	ar3_1 = ar1

	fmt.Println("ar3_1", ar3_1)

	/* cannot use ar1 (type [4]int) as type [5]int in assignment
	var ar3_2 [5]int
	ar3_2 = ar1
	*/


	// ar4 and ar5 are not equal or same
	var ar4 [4]int
	var ar5 [5]int
	fmt.Println("is ar4 equal to ar5: ", reflect.DeepEqual(ar4, ar5))
}


func itemExists(arrayType interface{}, item interface{}) bool {

	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invlid data type")
	}

	for i :=0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}



