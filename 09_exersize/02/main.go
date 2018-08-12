package main

import (
	"fmt"
)

func main() {

	var smallint int
	var largeint int
	fmt.Println("Enter Small Number:")
	fmt.Scan(&smallint)
	fmt.Println("Enter Large Number:")
	fmt.Scan(&largeint)

	remainder := largeint % smallint
	fmt.Println("remainder:", remainder)
}

/*
Create a program that prints to the terminal asking for a user to enter a
small number and a larger number. Print the remainder of the larger
number divided by the smaller
*/
