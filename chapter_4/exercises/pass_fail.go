package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {

	fmt.Print("Enter a grage: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failed"
	}
	fmt.Println("A grade of", grade, "is", status)

}
