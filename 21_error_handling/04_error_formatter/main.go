package main

import (
	"fmt"
	"log"
)

func main() {
	sq, err := sqrt(-5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sq)
}

func sqrt(n int) (int, error) {
	if n < 0 {
		var ErrNegativeNumber = fmt.Errorf("Invalid Parm: Negative Number %v", n)
		return 0, ErrNegativeNumber
	}
	return n * n, nil
}
