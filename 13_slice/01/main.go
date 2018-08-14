package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 1)
	mySlice[0] = 50
	fmt.Println(mySlice)
	fmt.Println(mySlice[0])
	//mySlice[1] = 10  //panic: runtime error: index out of range
	mySlice = append(mySlice, 20)
	fmt.Println(mySlice)
}
