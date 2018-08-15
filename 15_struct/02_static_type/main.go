package main

import (
	"fmt"
)

type foo int

func main() {
	var myInt foo

	fmt.Printf("%T %v \n", myInt, myInt)

	var newInt int

	fmt.Printf("%T %v \n", newInt, newInt)

	myInt = 50
	newInt = 100

	fmt.Println(int(myInt) + newInt)
}
