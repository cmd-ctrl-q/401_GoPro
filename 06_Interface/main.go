/*

	Interfaces.

	Everything is an interface! ie first class citizen

*/

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

type shape interface { // everything in go
	area() float64
}

// implements shape by using area() method from shape
func (s square) area() float64 {
	return s.side * s.side
}

// implements shape by using area() method from shape
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// simple function.
func info(x shape) {
	fmt.Println(x)
	fmt.Println(x.area())
}

func main() {

	sq := square{
		side: 10,
	}
	ci := circle{10}

	info(sq)
	info(ci)
}
