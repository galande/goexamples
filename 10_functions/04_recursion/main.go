package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(4))
}

func factorial(n uint) uint {
	if n == 0 {
		fmt.Println("n is:", n)
		return 1
	}
	fmt.Println("I AM before")
	return n * factorial(n-1)
}
