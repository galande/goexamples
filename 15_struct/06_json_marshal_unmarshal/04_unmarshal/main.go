package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	var p1 Person
	fmt.Println(p1)

	bs := []byte(`{"First":"Bhausaheb", "Last":"Galande", "Age":25}`)

	fmt.Println(bs)

	json.Unmarshal(bs, &p1)

	fmt.Println(p1)

}
