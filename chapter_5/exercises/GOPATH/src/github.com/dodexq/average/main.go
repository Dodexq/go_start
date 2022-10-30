// avarage вычесляет среднее значение
package main

import "fmt"

func main() {
	var numbers = [3]float64{71.2, 52.6, 84.2}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("%0.2f\n", sum/sampleCount)
}
