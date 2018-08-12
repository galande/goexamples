package main

import (
	"fmt"
)

func main() {

	greeting := func() {
		fmt.Println("hello")
	}

	greeting()
	fmt.Printf("%T \n", greeting)
}
