package main

import (
	"fmt"
	"log"

	"github.com/dodexq/keyboard"
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
