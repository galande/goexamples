package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radious float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}

func (s square) area() float64 {
	return s.side * s.side
}

func getInfo(sp shape) {
	fmt.Println("Area:", sp.area())
}

func main() {

	s := square{5}
	c := circle{5}

	getInfo(s)
	getInfo(c)
}
