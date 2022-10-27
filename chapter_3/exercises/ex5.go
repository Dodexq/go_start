package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(number *int) {
	*number *= 2
}

func main() {
	var myIntPointer *float64 = createPointer()
	fmt.Println(*myIntPointer)
	fmt.Println(myIntPointer)
	fmt.Println()
	//
	var myBool bool = true
	printPointer(&myBool)
	fmt.Println()
	//
	amount := 6
	double(&amount)
	fmt.Println(amount)

}
