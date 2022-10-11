package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello, GO")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	fmt.Println(5 > 2)

	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, GO!"))

	var quantity int
	var length, width float64
	var str string
	quantity = 5

	fmt.Println()
	str = "Hello, GO!"
	fmt.Println(str, quantity, "times")
	length, width = 5.2, 2.4
	fmt.Println(length, width)

	fmt.Println()
	var quantity2 int = 4
	var length2, width2 float64 = 6.1, 4.0
	var str2 string = "Go, Hello!"

	fmt.Println(reflect.TypeOf(quantity2))
	fmt.Println(reflect.TypeOf(length2))
	fmt.Println(reflect.TypeOf(width2))
	fmt.Println(reflect.TypeOf(str2))

}
