package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	age   int
}

func main() {
	p1 := Person{"Bhausaheb", "Galande", 25}

	bs, _ := json.Marshal(p1)

	fmt.Printf("%T \n", bs)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
