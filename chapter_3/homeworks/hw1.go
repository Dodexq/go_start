package main

import (
	"fmt"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get sqare root of negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squareRoot(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(root)
	}

	root2, err := squareRoot(-4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(root2)
	}

}
