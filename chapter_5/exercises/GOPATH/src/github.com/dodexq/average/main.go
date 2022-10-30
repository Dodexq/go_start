// avarage вычесляет среднее значение
package main

import (
	"fmt"
	"log"

	"github.com/dodexq/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("%0.2f\n", sum/sampleCount)
}
