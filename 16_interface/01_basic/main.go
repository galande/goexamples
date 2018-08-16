package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Printf("%T \n", z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
