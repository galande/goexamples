package main

import (
	"fmt"
)

func main() {

	myMap := make(map[string]string)

	fmt.Println(myMap)

	myMap["FirstName"] = "Bhausaheb"
	myMap["LastName"] = "Galande"

	fmt.Println(myMap)

	fmt.Println(myMap["LastName"])

	fmt.Println("Length Of Map:", len(myMap))

	value, isExist := myMap["FirstName"]

	fmt.Println(value, isExist)

	myMap["Designation"] = "Blockchain Developer"

	fmt.Println(myMap)

	delete(myMap, "LatName")

	fmt.Println(myMap)
}
