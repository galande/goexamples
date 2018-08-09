package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println("Before Change Value:", x)
	fmt.Println("Memory address of x:", &x)

	zero(&x)
	fmt.Println("After Change Value:", x)
}

func zero(z *int) {
	fmt.Println("z is:", z)
	fmt.Println("*z is", *z)
	*z = 0
}
