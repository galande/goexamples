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

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

func (p person) fullName() string {
	//fmt.Println(p)
	return p.first + " " + p.last
}
