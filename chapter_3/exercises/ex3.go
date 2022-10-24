package main

import (
	"errors"
	"fmt"
	"math"
)

func maniReturns() (int, bool, string) {
	return 1, true, "hello"

}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())

	err2 := fmt.Errorf("a height of %0.2f is invalid", -2.3333)
	fmt.Println(err2.Error())
	fmt.Println(err2)
	//
	fmt.Println()
	myInt, myBool, myString := maniReturns()
	fmt.Println(myInt, myBool, myString)
	//
	fmt.Println()
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)
}
