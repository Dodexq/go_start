package main

import "fmt"

func main() {

	// Сокращенное объявление переменных с присвоением
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println(length*width, "sqare meters")

	// Преобразования переменных
	var length2 float64 = 1.2
	var width2 int = 2

	fmt.Println("Area is", length2*float64(width2))
	fmt.Println("length2 > width2?", length2 > float64(width2))

}
