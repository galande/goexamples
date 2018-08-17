package main

import (
	"fmt"
)

func main() {

	for n := range genFactorial(factorial(5, 10, 12, 13, 15)) {
		fmt.Println(n)
	}
}

func factorial(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func genFactorial(c chan int) chan int {

	out := make(chan int)

	go func() {
		for n := range c {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()

	return out
}
