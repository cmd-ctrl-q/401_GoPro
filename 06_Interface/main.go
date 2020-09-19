/*

	*** Interfaces ***

	Implements
	- a type implements an some interface if it has a method with the exact signature of it. 

*/

package main

import (
	"fmt"
	"math"
)

// square is a pointer to a struct value
type square struct {
	side float64
}

// constructor-like function to prevent nil or zero values
func New(side float64) square {
	sq := square{side}
	return sq
}

type circle struct {
	radius float64
}

type shape interface { 
	area() float64
}

// implements shape by using the signature area() from shape
func (s *square) area() float64 { // takes in pointer
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
	sq2 := New(42) // using a constructor-like function
	ci := circle{10}

	info(&sq) // generate's pointer to the operand. return &{10} shows struct is a pointer
	info(&sq2)
	info(ci)
}
