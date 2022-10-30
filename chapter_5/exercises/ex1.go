package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[4])
	fmt.Println(reflect.TypeOf(notes[0]))
	fmt.Println(reflect.TypeOf(notes))
	fmt.Println()
	//
	var primes [5]int
	primes[0] = 5
	primes[1] = 14
	fmt.Println(primes[1])
	//
	var dates [3]time.Time
	dates[0] = time.Unix(1237614000, 0)
	dates[1] = time.Unix(1477220000, 0)
	dates[2] = time.Unix(1208632200, 0)
	fmt.Println(dates[0])
	fmt.Println(dates[1].Date())
	fmt.Println()
	//
	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])
	fmt.Println()
	//
	var notes1 [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes1[0], notes1[3], notes1[6])
	var primes1 [5]int = [5]int{2, 3, 19, 52, 64}
	fmt.Println(primes1[4])
	fmt.Println()
	//
	var notes3 [3]string = [3]string{"one", "two", "three"}
	var primes3 [5]int = [5]int{2, 3, 19, 52, 64}
	fmt.Println(notes3)
	fmt.Println(primes3)

	fmt.Printf("%#v\n", notes3)
	fmt.Printf("%#v\n", primes3)
	fmt.Println()
	//
	notes = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
	fmt.Println()
	//
	fmt.Println(len(notes), "- massive len")
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	fmt.Println()
	//

}
