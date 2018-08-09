package main

import (
	"fmt"
)

func main() {
	var x = 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment()) //print 1
	fmt.Println(increment()) //print 2
}
