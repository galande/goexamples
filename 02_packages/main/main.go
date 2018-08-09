package main

import (
	"fmt"

	"github.com/galande/udemy/goexamples/02_packages/stringutils"
)

func main() {
	fmt.Println("Main Package")
	fmt.Println(stringutils.Myname)
	fmt.Println(stringutils.Getstring("Galande"))
}
