package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes), "- massive len")
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	fmt.Println()
	//
	for index, note := range notes {
		fmt.Println(index, note)
	}
	fmt.Println()
	//

}
