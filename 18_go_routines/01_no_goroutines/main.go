package main

import (
	"fmt"
)

func main() {
	//Runs one by one
	foo()
	bar()
}

func foo() {
	for i := 0; i <= 20; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i <= 20; i++ {
		fmt.Println("bar:", i)
	}
}
