package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{"Bhausaheb", "Galande", 25}
	p2 := person{"Tanmay", "Khambekar", 26}

	fmt.Println(fullName(p1))
	fmt.Println(fullName(p2))
}

func fullName(p person) string {
	return p.first + " " + p.last
}
