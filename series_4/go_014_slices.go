package main

import "fmt"

func main() {

	sliceBasics()
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

	fmt.Printf(" s0 : %#v \n", s0)
	fmt.Printf(" s1 : %#v \n", s1)
	fmt.Printf(" s2 : %#v \n", s2)
	fmt.Printf(" s3 : %#v \n", s3)

}