package main

import (
	"fmt"
)

func main() {

	type person struct {
		first  string
		second string
		age    int
	}

	p1 := person{"Bhausaheb", "Galande", 25}

	fmt.Println(p1.first, p1.second, p1.age)
}
