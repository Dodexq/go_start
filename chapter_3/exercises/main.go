package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0

}

func main() {
	var amount, total float64
	amount = paintNeeded(5.1, 3.2)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount

	amount = paintNeeded(3.8, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount

	fmt.Printf("Total: %0.2f liters\n", total)

}
