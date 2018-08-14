package main

import (
	"fmt"
)

func main() {
	var alpha [58]string

	fmt.Println(alpha)
	for i := 65; i < 123; i++ {
		alpha[i-65] = string(i)
	}

	fmt.Println(alpha)
}
