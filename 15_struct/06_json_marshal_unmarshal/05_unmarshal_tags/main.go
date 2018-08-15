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

	var p1 Person
	fmt.Println(p1)

	bs := []byte(`{"First":"Bhausaheb", "Last":"Galande", "Years Since":25}`)

	fmt.Println(bs)

	json.Unmarshal(bs, &p1)

	fmt.Println(p1)

}
