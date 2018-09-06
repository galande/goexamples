package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNegativeNumber = errors.New("Invalid Parm: Negative Number!!")

func main() {
	sq, err := sqrt(-5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sq)
}

func sqrt(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeNumber
	}
	return n * n, nil
}
