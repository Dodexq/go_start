package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}

func status(grade float64) string {
	if grade < 60.0 {
		return "failing"
	}
	return "passing"
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func double(number float64) float64 {
	return number * 2
}

func main() {
	sayHi()
	repeatLine("hello", 3)
	println()
	//
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(5))
	println()
	//
	fmt.Println(status(72.0))
	fmt.Println(status(34.3))

}
