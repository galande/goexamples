package main

import (
	"fmt"
)

func main() {

	switch "blockchain" {
	case "AI":
		fmt.Println("artificial intelligence")
	case "RPA":
		fmt.Println("Robot Processing Automation")
	case "blockchain":
		fmt.Println("Blockchain!!..New Era of Internet")
		fallthrough //makes next case true
	case "ML":
		fmt.Println("Machine Leraning")
		fallthrough
	default:
		fmt.Println("Default IT")
	}
}
