package main

import (
	"fmt"
)

func main() {

	switch "Blockchain" {
	case "AI", "ML":
		fmt.Println("artificial intelligence", "Machine Leraning")
	case "RPA", "Blockchain":
		fmt.Println("Robot Processing Automation", "Blockchain")
	default:
		fmt.Println("Default IT")
	}
}
