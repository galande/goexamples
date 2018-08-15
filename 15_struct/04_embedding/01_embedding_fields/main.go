package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type armyPerson struct {
	person
	licenceToKill bool
}

func main() {

	p1 := person{"Bhausaheb", "Galande", 25}
	p2 := person{"Tanmay", "Khambekar", 26}

	p3 := armyPerson{
		person:        person{"Sandip", "Patil", 24},
		licenceToKill: true,
	}

	p4 := armyPerson{p2, true}

	fmt.Println(p1)
	fmt.Println(p3)
	fmt.Println(p4)
}
