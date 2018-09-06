package main

import (
	"fmt"
)

func main() {
	fmt.Println("I am in main function")
}

func init() {
	fmt.Println("I am in init. Will run first..")
}
