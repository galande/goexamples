package main

import (
	"errors"
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
		return 0, errors.New("Sqaure of negative number!!!!")
	}
	return n * n, nil
}
