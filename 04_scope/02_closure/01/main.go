package main

import "fmt"

//package level scope
var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment()) // print 1
	fmt.Println(increment()) // print 2
}
