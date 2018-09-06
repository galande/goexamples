package main

import (
	"fmt"
)

func main() {

	c1 := incrementor()
	//c2 := incrementor()
	c2 := createCounter(c1)

	for total := range c2 {
		fmt.Println("******************")
		fmt.Println(total)
	}

}

func incrementor() chan int {

	out := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func createCounter(c chan int) chan int {
	out := make(chan int)
	go func() {
		var total int
		for n := range c {
			fmt.Println(n)
			total++
		}
		out <- total
		close(out)
	}()

	return out
}
