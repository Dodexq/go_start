package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
	// or
	today := time.Now().Day()
	fmt.Println(today)
	fmt.Println(time.Now().Date())
}
