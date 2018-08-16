package main

import (
	"fmt"
)

func main() {

	go foo()
	go bar()
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

//Nothing will print because different rotines runs differently.
