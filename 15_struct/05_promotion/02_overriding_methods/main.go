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

	p1 := person{"Sandip", "Patil", 24}

	p2 := armyPerson{
		person:        person{"Sagar", "Tadas", 24},
		first:         "army first",
		licenceToKill: true,
	}

	p1.greetings()
	p2.greetings()
	p2.person.greetings()

}

func (p person) greetings() {
	fmt.Println(p.first, " Normal person.")
}

func (ap armyPerson) greetings() {
	fmt.Println(ap.first, " special army person")
}
