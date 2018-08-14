package main

import (
	"fmt"
)

func main() {
	mySlice := make([]string, 3, 5)
	fmt.Println("Length:", len(mySlice))
	fmt.Println("Capacity", cap(mySlice))

	mySlice[0] = "Bhausaheb"
	mySlice[1] = "Vijay"
	mySlice[2] = "DJ"

	mySlice = append(mySlice, "Kiran")
	fmt.Println("Length:", len(mySlice))
	fmt.Println("Capacity", cap(mySlice))

	mySlice = append(mySlice, "Danny")
	fmt.Println("Length:", len(mySlice))
	fmt.Println("Capacity", cap(mySlice))

	//Here length is 5 and capacity is also 5.
	//After full capacity , if we append then capacity will automatically incresed
	//(Doubled). i.e-> New underline array will be created with double capacity and
	//old values will be copied there.

	mySlice = append(mySlice, "Sagar")
	fmt.Println("Length:", len(mySlice))
	fmt.Println("Capacity", cap(mySlice))

	fmt.Println(mySlice)
}
