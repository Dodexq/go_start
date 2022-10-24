package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil

}

func main() {
	// var amount, total float64
	// amount = paintNeeded(5.1, 3.2)
	// fmt.Printf("%0.2f liters needed\n", amount)
	// total += amount

	// amount = paintNeeded(3.8, 3.5)
	// fmt.Printf("%0.2f liters needed\n", amount)
	// total += amount

	// fmt.Printf("Total: %0.2f liters\n", total)

	amount, err := paintNeeded(3.4, 3.5)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}

	amount2, err := paintNeeded(3.4, -2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount2)
	}

}
