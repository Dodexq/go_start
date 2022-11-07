package main

import (
	"fmt"
	"log"

	"github.com/dodexq/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
