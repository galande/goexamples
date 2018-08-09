package main

import (
	"fmt"
)

func swapper() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func main() {

	increment := swapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
