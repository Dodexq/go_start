package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 32, "Rohit": 52.3, "Carl": 86.5}
	for name, grade := range grades {
		fmt.Printf("%s has a grade %0.1f%%\n", name, grade)
	}
	fmt.Println("^ random sort")
	fmt.Println()
	//
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}

}
