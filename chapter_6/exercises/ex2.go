package main

import "fmt"

func main() {
	myArray := [5]string{"a", "b", "c", "d", "e"}
	mySlice := myArray[0:3]
	fmt.Println(mySlice)
	i, j := 1, 4
	fmt.Println(myArray[i:j])
	mySlice2 := myArray[2:]
	fmt.Println(mySlice2)
	//
	myArray3 := [5]string{"a", "b", "c", "d", "e"}
	mySlice3 := myArray3[0:3]
	mySlice4 := myArray3[2:5]
	fmt.Println(mySlice3, mySlice4)
	//
	myArray4 := [5]string{"a", "b", "c", "d", "e"}
	mySlice5 := myArray4[0:3]
	myArray4[1] = "X"
	fmt.Println(myArray4)
	fmt.Println(mySlice5)
	fmt.Println()
	//
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))
	fmt.Println()
	//
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "append"
	fmt.Println(s1, s2, s3, s4)
	// better
	s1 = []string{"s1", "s1"}
	s1 = append(s1, "s2", "s2")
	s1 = append(s1, "s3", "s3")
	s1 = append(s1, "s4", "s4")
	fmt.Println(s1)
	fmt.Println()
	//
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[5], boolSlice[7])
	//
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)
	fmt.Println()
	//
	var mySlice6 []string
	if len(mySlice6) == 0 {
		mySlice6 = append(mySlice6, "first item")
	}
	fmt.Printf("%#v\n", mySlice6)

}
