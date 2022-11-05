package main

import "fmt"

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func main() {
	severalInts(1)
	severalInts(1, 2, 3)
	fmt.Println()
	//
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()
}
