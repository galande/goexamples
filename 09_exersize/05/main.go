package main

import (
	"fmt"
)

func main() {

	result := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			result++
		} else if i%5 == 0 {
			result++
		}
	}
	fmt.Println("Result:", result)
}
