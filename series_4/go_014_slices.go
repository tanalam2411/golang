package main

import "fmt"

func main() {

	fmt.Println("******************* sliceBasics ***********************")
	sliceBasics()
	fmt.Println("******************* sliceUsingMake ***********************")
	sliceUsingMake()
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
	var s2 = make([]int, 5, 10)

	fmt.Printf("s0: %#v, cap(s0): %#v, len(s0): %#v \n", s0, cap(s0), len(s0))
	fmt.Printf("s1: %#v, cap(s1): %#v, len(s1): %#v \n", s1, cap(s1), len(s1))
	fmt.Printf("s2: %#v, cap(s2): %#v, len(s2): %#v \n", s2, cap(s2), len(s2))

	s1[4] = 100
	s2[4] = 200
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

	for i :=0 ; i < len(s2) ; i++ {
		s2[i] = i + 10
	}

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
