package main

import (
	"fmt"
)

func main() {
	result := average(10, 12, 14, 16, 18, 20)
	fmt.Println(result)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
		//fmt.Println("v is:", v, "sf is:", sf)
	}

	return total / float64(len(sf))
}
