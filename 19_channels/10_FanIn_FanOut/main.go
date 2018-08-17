package main

import (
	"fmt"
	"sync"
)

func main() {
	c := gen(3, 5)

	c1 := sq(c)
	c2 := sq(c)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(c chan int) chan int {

	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(ch ...chan int) chan int {

	var wg sync.WaitGroup
	wg.Add(len(ch))

	out := make(chan int)
	for _, c := range ch {
		go func(c chan int) {
			out <- <-c
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
