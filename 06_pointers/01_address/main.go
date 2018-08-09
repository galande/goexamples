package main

import (
	"fmt"
)

const metertoyard = 1.09361

func main() {
	var meters float64
	fmt.Printf("Enter Meter:")
	fmt.Scan(&meters)
	yard := meters * metertoyard
	fmt.Println("Meter is", meters, "and Yard is", yard)
}
