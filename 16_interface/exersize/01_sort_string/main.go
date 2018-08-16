package main

import (
	"fmt"
	"sort"
)

func main() {

	type person []string

	studyGroup := person{"Bhausaheb", "Vijay", "Dattatray", "Kuldeep"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
}
