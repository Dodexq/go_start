package main

import "fmt"

func sum(numbers ...int) int {
	var sum int = 0
	for _, nunumber := range numbers {
		sum += nunumber
	}
	return sum
}

func main() {
	fmt.Println(sum(9, 7))
	fmt.Println(sum(1, 2, 4))
}
