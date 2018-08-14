package main

import (
	"fmt"
)

func main() {
	firstInt := []int{1, 2, 3, 4, 5}
	secondInt := []int{6, 7, 8, 9, 10}

	firstInt = append(firstInt, secondInt...)
	fmt.Println(firstInt)
}
