package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for index := 0; index < 10; index++ {
			c <- index
		}
		fmt.Println("I am at close")
		close(c)
	}()

	//Doing something till channel does.
	fmt.Println("I am before range")

	for n := range c {
		fmt.Println(n)
	}

	fmt.Println("I am after range")
}
