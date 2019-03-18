package main

import (
	"fmt"
)

/*
A type determines the set of values and operations specific to values of that type.
Types may be named or unnamed.
Named types are specified by a(possibly qualified) type name
Unnamed types are specified using a type literal, which composes a new type from existing types.
*/

/* func1
Type Identity

- Two types are either identical or different
- Two named types are identical if their type names originate in the same TypeSpec.
  A named and an unamed type are always different.
  Two unnamed types are identical if the corresponding type literal are identical,
  that is, if they have the same literal structure and corresponding components have identical types.
*/

/* func2
Assignability

A value x is assignable to a variable of type T("x is assignable to T") in any these cases:
- x's type is identical to T
- x's type V and T have identical underlying types and atleast one of V or T  is not a named type
- x is a bidirectional channel value, T is a channel type, x's type V and T have identical element types
  and atleast one of V or T is not a named type.
- x is predeclared identifier nil and T is a pointer, function, slice, map, channel or interface type.
- x is an untyped constant representable by a value of type T.
*/

// named type StudId
type StudId int

func main() {

	stud0 := StudId(25) // var stud0 StudId = 25
	printStudId(stud0)
	printStudId(30)

	//id := 35
	//printStudId(id) //  cannot use id (type int) as type StudId in argument to printStudId
	// This is simlilar to time.duration
	//time.Sleep(10 * time.Nanosecond)
	//t := 20
	//time.Sleep(t * time.Nanosecond) //  t * time.Nanosecond (mismatched types int and time.Duration)
	// time.go // type Duration int64
	// time.Sleep(time.Duration(t) * time.Nanosecond)

	fmt.Println("---------------- func1 ---------------------")
	func1()

	fmt.Println("---------------- func2 ---------------------")
	func2()

	fmt.Println("---------------- func3 ---------------------")
	func3()

	fmt.Println("---------------- func4 ---------------------")
	func4()
}

func printStudId(id StudId) {
	fmt.Printf("Id : %T -> %v \n", id, id)
}

type DataSet []int
type DataSet2 []int
type Dict map[string]string


func func1() {

	var data0 = DataSet([]int{1, 2, 3, 4, 5})
	var data1 = []int{6, 7, 8, 9, 10}
	var datax = []int{}
	var data2 = DataSet2([]int{11, 12, 13, 14, 15})

	fmt.Printf("data0: %T -> %v\n", data0, data0)
	fmt.Printf("data1: %T -> %v\n", data1, data1)
	fmt.Printf("datax: %T -> %v\n", datax, datax)
	fmt.Printf("data2: %T -> %v\n", data2, data2)


	var m0 map[string]string
	var m1 Dict
	fmt.Printf("m0: %T -> %v\n", m0, m0)
	fmt.Printf("m1: %T -> %v\n", m1, m1)


	var p0 struct{
		n string
		a int
		ssn string
	}

	var p1 struct{
		n string
		a int
		ssn string
	}

	fmt.Printf("p0: %T -> %v\n", p0, p0)
	fmt.Printf("p1: %T -> %v\n", p1, p1)
	// p0 = p1 // will work, but if any order of attributes is changed than p0 = p1 will not be similar and throw error

}


type V []int
type T []int

func func2() {
	var v = []int{1, 2, 3}
	var t T = T([]int{4, 5, 6})
	var t1 = T([]int{7, 8, 9})
	t = v

	fmt.Println(v, t)

	v = t1
	fmt.Println(v, t)
}


type C chan int

func func3() {
	x0 := make(chan int)

	var c C
	c = x0

	var in chan<- int
	var out <-chan int

	in = c
	out = c

	fmt.Println(in)
	fmt.Println(out)
	fmt.Println(c)
}


const Con = 0x929399399494
type Ch chan int

func func4() {

	var p *int
	var f func(int) string

	var m map[string]float32
	var s []complex64
	var ch chan T

	p = nil
	f = nil
	m = nil
	s = nil
	ch = nil

	var i float64 = Con
	fmt.Println(p, f, m, s, ch, i)

}


/* output

Id : main.StudId -> 25
Id : main.StudId -> 30
---------------- func1 ---------------------
data0: main.DataSet -> [1 2 3 4 5]
data1: []int -> [6 7 8 9 10]
datax: []int -> []
data2: main.DataSet2 -> [11 12 13 14 15]
m0: map[string]string -> map[]
m1: main.Dict -> map[]
p0: struct { n string; a int; ssn string } -> { 0 }
p1: struct { n string; a int; ssn string } -> { 0 }
---------------- func2 ---------------------
[1 2 3] [1 2 3]
[7 8 9] [1 2 3]
---------------- func3 ---------------------
0xc00001e0c0
0xc00001e0c0
0xc00001e0c0
---------------- func4 ---------------------
<nil> <nil> map[] [] <nil> 1.61162628535444e+14

*/