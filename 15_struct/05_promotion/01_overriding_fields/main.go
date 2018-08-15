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
	first         string
	licenceToKill bool
}

func main() {

	p1 := armyPerson{
		person:        person{"Sandip", "Patil", 24},
		first:         "army first",
		licenceToKill: true,
	}

	fmt.Println(p1.first)
	fmt.Println(p1.person.first)
}
