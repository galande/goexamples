package main

import (
	"fmt"
)

func main() {
	// %x for lowercase hex values
	fmt.Printf("%x \n", 42)

	// %X for uppercase hex values
	fmt.Printf("%X \n", 42)

	// # is to add 0x in hex values
	fmt.Printf("%#x \n", 42)
	fmt.Printf("%#X \n", 42)
}
