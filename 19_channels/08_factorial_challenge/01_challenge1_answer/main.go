package main

import (
	"fmt"
)

func main() {

	c := factorial(5)
	fmt.Println("Result:", <-c)
}

func factorial(n int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}
