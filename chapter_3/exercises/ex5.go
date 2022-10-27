package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println()
	fmt.Println(*myBoolPointer)
}

func main() {
	var myIntPointer *float64 = createPointer()
	fmt.Println(*myIntPointer)
	fmt.Println(myIntPointer)
	//
	var myBool bool = true
	printPointer(&myBool)
}
