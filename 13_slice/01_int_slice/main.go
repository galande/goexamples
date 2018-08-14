package main

import (
	"fmt"
)

func main() {

	myInt := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T \n", myInt)
	fmt.Println(len(myInt))
	fmt.Println(cap(myInt))
	fmt.Println(myInt)

	myInt1 := []int{}
	fmt.Println(len(myInt1))
	fmt.Println(cap(myInt1))
	fmt.Println(myInt1)

}
