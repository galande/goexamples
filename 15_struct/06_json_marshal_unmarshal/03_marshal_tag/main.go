package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"Years Since"`
}

func main() {
	p1 := Person{"Bhausaheb", "Galande", 25}

	fmt.Println(p1)

	bs, _ := json.Marshal(p1)
	fmt.Println(bs)

	fmt.Println(string(bs))
}
