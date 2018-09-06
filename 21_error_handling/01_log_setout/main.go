package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	fp, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	log.SetOutput(fp)
}

func main() {
	_, err := os.Open("abcd.txt")
	if err != nil {
		log.Println("Error Occured:", err)
		log.Println("Error Occured:", err)
	}
}
