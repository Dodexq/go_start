package main

import "fmt"

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paintNeeded(5.2, 2.8)
	paintNeeded(3.2, 3.5)
	paintNeeded(6.4, 3.3)
}
