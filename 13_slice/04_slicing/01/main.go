package main

import (
	"fmt"
)

func main() {

	var result []int
	fmt.Println(len(result))
	fmt.Println(cap(result))
	fmt.Println(result)

	myString := []string{"A", "B", "C", "D", "E", "F", "G"}
	fmt.Println(myString)
	fmt.Println(myString[2:4])
	fmt.Println(myString[1])
	fmt.Println(myString[:5])
	fmt.Println(myString[2:])
	fmt.Println("BHAU"[2]) //65 > A UTF-8 65
}
