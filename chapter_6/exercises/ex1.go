package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println()
	//
	notes1 := make([]string, 6)
	primes1 := make([]int, 10)
	fmt.Println(len(notes1))
	fmt.Println(len(primes1))
	fmt.Println()
	//
	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}
	fmt.Println()
	//
	notes2 := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes2[2], notes2[4], notes2[6])
	primes2 := []int{
		5,
		15,
		25,
	}
	fmt.Println(primes2[0], primes2[1], primes2[0]+primes2[1])
}
