package main

import (
	"fmt"
)

func main() {

	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

func visit(val []int, callback func(n int)) {
	for _, n := range val {
		callback(n)
	}
}
