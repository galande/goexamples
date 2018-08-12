package main

import (
	"fmt"
)

func main() {

	fmt.Println(findMax(1, 2, 3, 4, 35, 5))
}

func findMax(numbers ...uint) uint {
	var max uint
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

/*

Write a function with one variadic parameter that finds the greatest
number in a list of numbers.

*/
