package main

import (
	"fmt"
)

func main() {
	defer world()
	hello()
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}
