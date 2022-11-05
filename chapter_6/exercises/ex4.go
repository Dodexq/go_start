package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(maximum(41.8, 57.6, 79.2))
	fmt.Println(maximum(91.8, 67.6, 79.2, 92.6))
	fmt.Println()
	//
	fmt.Println(inRange(0, 100, 12, 41, 512, 69, 11, -12, 51, 112, 220))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}
