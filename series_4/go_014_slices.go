package main

import "fmt"

func main() {

	fmt.Println("******************* sliceBasics ***********************")
	sliceBasics()

	fmt.Println("******************* sliceUsingMake ***********************")
	sliceUsingMake()

	fmt.Println("******************* sliceAppend ***********************")
	sliceAppend()

	fmt.Println("******************* spreadOperator ***********************")
	spreadOperator()

	fmt.Println("******************* copySlices ***********************")
	copySlices()

	fmt.Println("******************* iterRangeOnlyIndex ***********************")
	iterRangeOnlyIndex()
}

func sliceBasics() {
	var nums = [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Printf(" nums : %#v \n", nums)

	var mid = len(nums) / 2
	fmt.Printf("mid: %#v, %T \n", mid, mid)

	var s0 = nums[:]
	var s1 = nums[:mid]
	var s2 = nums[mid:]
	var s3 = nums[3:7]

	// capacity is - len from start offset of slice till end of the underlying array
	fmt.Printf(" s0 [0:11] : %#v cap: %#v\n", s0, cap(s0))
	fmt.Printf(" s1 [0:5] : %#v cap: %#v\n", s1, cap(s1))
	fmt.Printf(" s2 [5:11] : %#v cap: %#v\n", s2, cap(s2))
	fmt.Printf(" s3 [3:6] : %#v cap: %#v\n", s3, cap(s3))

	//fmt.Printf("s3[cap(s3)]", s3[cap(s3)]) // panic: runtime error: index out of range

	// even slice s3 was - []int{13, 14, 15, 16}, when we try to pull more elements it fetches with in the underlying array
	fmt.Printf("s3[:cap(s3)] : %#v \n", s3[:cap(s3)])

}


func sliceUsingMake() {

	var s0 = []int{}
	var s1 = make([]int, 5)
	var s2 = make([]int, 0, 10)

	fmt.Printf("s0: %#v, cap(s0): %#v, len(s0): %#v \n", s0, cap(s0), len(s0))
	fmt.Printf("s1: %#v, cap(s1): %#v, len(s1): %#v \n", s1, cap(s1), len(s1))
	fmt.Printf("s2: %#v, cap(s2): %#v, len(s2): %#v \n", s2, cap(s2), len(s2))

	s1[4] = 100
	//s2[4] = 200
	// s2[9] = 200 // panic: runtime error: index out of range
	fmt.Printf("s1[10]: %#v \n", s1)
	fmt.Printf("s2[10]: %#v \n", s2)

	for i, v := range s1 {
		fmt.Println("s1:", i, v)
	}

	for i, v := range s2 {
		fmt.Println("s2:", i, v)
	}

	for i :=0 ; i < cap(s1) ; i++ {
		s1[i] = i + 10
	}

	for i :=0 ; i < cap(s2) ; i++ {
		//s2[i] = i + 10
		s2 = append(s2, i + 10)
	}


	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	var testScores []int

	var count = 10

	testScores = make([]int, count)
	testScores[0] = 56
	testScores[1] = 67
	testScores[2] = 43

	printSlice("testScores", testScores)
}


func printSlice(n string, a []int) {
	fmt.Printf("%v = %#v, len = %v, cap = %v \n", n, a, len(a), cap(a))
}


func sliceAppend() {
	var s0 []int
	printSlice("s0", s0)

	s0 = make([]int, 2, 3)  // len -2 , cap - 3
	printSlice("s0", s0)

	s0[0] = 11
	s0[1] = 12
	printSlice("s0", s0)

	//s0 = append(s0, 13, 14, 15)
	s0 = appendImpl(s0, 13, 14, 15)
	printSlice("s0", s0)
}


func appendImpl(s []int, e ...int) []int {
	fmt.Println("s len", len(s))
	fmt.Println("s cap", cap(s))
	fmt.Println("e len", len(e))
	fmt.Println("e cap", cap(e))

	var t []int
	if len(e) > (cap(s)-len(s)) {
		t = make([]int, cap(s)+len(e))
		for i:=0; i < len(s); i++{
			t[i] = s[i]
		}

	} else {
		t = s[:len(s)+len(e)]
	}

	fmt.Println("---", len(e), cap(e), e, "---", len(s), cap(s), s)
	fmt.Println(t, len(t), cap(t))

	for j:=0; j<len(e) ; j++ {
		t[len(s)+j] = e[j]
	}

	return t
}


func spreadOperator() {

	var a1 = []int{11, 22, 33, 44, 55}
	var s1 = []int{110, 220, 330, 440, 550}

	variadicFunction(a1...) // variadicFunction(a1[0], a1[1], a1[2], a1[3], a1[4])
	variadicFunction(s1...) // variadicFunction(s1[0], s1[1], s1[2], s1[3], s1[4])
}


func variadicFunction(a ...int) {
	for _, i := range a {
		fmt.Println(i)
	}
}


func copySlices() {
	
	var dst []int
	copy(dst, []int{12, 23})
	printSlice("dst", dst)
	
	var s = []int{4, 56}
	dst = make([]int, 10)
	copy(dst, s)
	printSlice("dst", dst)

	var t = []int{40, 563}
	dst = make([]int, 2)
	copy(dst, t)
	printSlice("dst", dst)

	s = []int{1, 2, 3, 4, 5}
	printSlice("s", s)

	var src = s[:4]
	dst = s[1:4]
	printSlice("Before copy src", src)
	printSlice("Before copy dst", dst)

	fmt.Println("copying - copy(dst, src)")
	copy(dst, src)
	printSlice("After copy src", src)
	printSlice("After copy dst", dst)
	printSlice("After copy s", s)


}


func iterRangeOnlyIndex() {

	s1 := []int{10, 20, 30, 40, 50}


	for index, value := range(s1) {
		fmt.Println(index, value)
	}


	for index, _ := range(s1) {
		fmt.Println(index)
	}

	for _, value := range(s1) {
		fmt.Println(value)
	}

	for index := range(s1) {
		fmt.Println(index)
	}
}