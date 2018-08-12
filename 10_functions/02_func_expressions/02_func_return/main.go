package main

import (
	"fmt"
)

func main() {

	greet := returnGreet()

	fmt.Println(greet())
	fmt.Printf("%T ", greet)

}

func returnGreet() func() string {

	return func() string {
		return "Hello..Hyperledger Fabric"
	}
}
