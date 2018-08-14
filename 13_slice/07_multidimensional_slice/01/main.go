package main

import (
	"fmt"
)

func main() {
	var records [][]string

	student1 := make([]string, 2)
	student1[0] = "Bhausaheb"
	student1[1] = "Galande"

	student2 := make([]string, 2)
	student2[0] = "Dattatray"
	student2[1] = "Jadhav"

	records = append(records, student1)
	fmt.Println(records)

	records = append(records, student2)
	fmt.Println(records)

	fmt.Printf("%T \n", records)

}
