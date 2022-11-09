package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	} else {
		fmt.Printf("%s is pass!\n", name)
	}
}

func main() {
	status("Alma")
	status("Rohit")
	status("Carl")
	fmt.Println()
	//
	counters := map[string]int{"a": 5, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	_, ok = counters["c"]
	fmt.Println(ok)
	fmt.Println()
	//
	status("Alma")
	status("Rohit")
	status("Carl")
}
