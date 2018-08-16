package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 20; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 20; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("bar:", i)
	}
	wg.Done()
}
