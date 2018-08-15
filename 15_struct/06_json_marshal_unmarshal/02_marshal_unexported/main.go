package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"Bhausaheb", "Galande", 25}
	fmt.Println(p1)

	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
