package main

import (
	"fmt"
)

func main() {

	xs := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}

func filter(numbers []int, callback func(n int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}
