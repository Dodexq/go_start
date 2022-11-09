package main

import (
	"fmt"
)

func main() {
	var myMap map[string]float64
	//or
	var ranks map[string]int
	ranks = make(map[string]int)
	//or
	ranksShort := make(map[string]int)
	fmt.Println(myMap, ranks, ranksShort)
	fmt.Println()
	//
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	fmt.Println()
	//
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
	fmt.Println()
	//
	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4], isPrime[7])
	fmt.Println()
	//
	ranks2 := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranks2["bronze"], ranks2["gold"])
	elements2 := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(elements2["H"])
	fmt.Println(elements2["Li"])
	fmt.Println()
	//
	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
}
