package main

import (
	"fmt"
	"log"

	"github.com/dodexq/keyboard"
)

func celsium(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {

	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsium := celsium(fahrenheit)
	fmt.Printf("%0.2f degrees Celsium\n", celsium)
}
