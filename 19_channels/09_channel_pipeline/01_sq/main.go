package main

import (
	"fmt"
)

func main() {
	c1 := generate(2, 3)
	c2 := square(c1)

	fmt.Println(<-c2)
	fmt.Println(<-c2)
}

func generate(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
			fmt.Println(n)
		}
		close(out)
	}()

	return out
}

func square(c chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()

	return out
}
