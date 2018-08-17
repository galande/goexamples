package main

import (
	"fmt"
)

func main() {
	f := factorial(5)
	fmt.Println("result:", f)
}

func factorial(n int) int {

	total := 1

	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

//Challenge Is:
// Use channel to get factorial
