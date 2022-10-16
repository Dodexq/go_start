package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("hw1.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
