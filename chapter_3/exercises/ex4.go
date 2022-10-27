package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myStr string
	fmt.Println(reflect.TypeOf(&myStr))
	fmt.Println()
	//
	var myInt2 int
	var myIntPointer *int
	myIntPointer = &myInt2
	fmt.Println(myIntPointer)

	var myFloat2 float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat2
	fmt.Println(myFloatPointer)
	fmt.Println()
	//
	myInt3 := 4
	myIntPointer3 := &myInt3
	fmt.Println(myIntPointer3)
	fmt.Println(*myIntPointer3)

	myFloat3 := 65.2
	myFloatPointer3 := &myFloat3
	fmt.Println(myFloatPointer3)
	fmt.Println(*myFloatPointer3)
	fmt.Println(reflect.TypeOf(myFloatPointer3))
	fmt.Println()
	//
	myInt4 := 23
	fmt.Println(myInt4)
	myIntPointer4 := &myInt4
	*myIntPointer4 = 10
	fmt.Println(*myIntPointer4)
	fmt.Println(myInt4)
}
