package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(reflect.TypeOf(replacer), "- type replacer variable")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed, "\n")

	// more shorter
	broken2 := "Hell# w#rld"
	fmt.Println(strings.NewReplacer("#", "o").Replace(broken2))
}
