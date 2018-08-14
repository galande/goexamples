package main

import (
	"fmt"
)

func main() {
	mySlice := make([]string, 3, 5)
	mySlice[0] = "Bhausaheb"
	mySlice[1] = "Vijay"
	mySlice[2] = "DJ"
	//mySlice[3] = "Kiran" //runtime error: index out of range

	fmt.Println(mySlice)
}
