package main

import "fmt"

func severalInts(numbers ...int) {
	for _, i := range numbers {
		fmt.Println(i)
	}
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	ints := []int{5, 12, 51, 52}
	severalInts(ints...)
	fmt.Println()
	stringSlice := []string{"a", "b", "c", "d", "e"}
	mix(5, true, stringSlice...)
}
