package main

import (
	"fmt"
)

func main() {

	x := 5
	fmt.Println("Before Change", x)
	zero(x)
	fmt.Println("After Change", x)
}

func zero(z int) {
	z = 0
}
