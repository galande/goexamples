package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}

//Only printing 0.
//Challege is , it should print 0 to 9.
