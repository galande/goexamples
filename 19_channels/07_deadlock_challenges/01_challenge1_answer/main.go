package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//Solution 1:- create goroutine  to send value into channel, so that it will be concurrent.

	go func() {
		c <- 1
		close(c)
	}()

	fmt.Println(<-c)

	//Solution2: Create goroutine to recieve value from channel

	// go func() {
	// 	fmt.Println(<-c)
	// }()

	// c <- 1
	// close(c)

}
