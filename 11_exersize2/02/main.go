package main

import (
	"fmt"
)

func main() {

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(half(1))
	fmt.Println(half(2))
}

/*
Modify the previous program to use a func expression.
*/
