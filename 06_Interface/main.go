/*

	*** Interfaces ***

	Struct embedding
		Embedding a struct in another struct to expose the embedded structs
		fields and methods for a struct to use.
	Implicit Interfacing / Mixins
		Struct that implement exact signature of an interface.
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// square is promoting its fields to cube
type cube struct {
	sq square // struct embedding
}

// New constructor-like function prevents nil or zero values
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

// Implicit Interfacing
// implements shape by using the signature area() from shape
// this can be viewed as method in a java class called Square.java
// only a pointer to a square can use this method
func (s *square) area() float64 { // takes in pointer
	return s.side * s.side
}

// implements shape by using area() method from shape
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// cube implements shape
func (c *cube) area() float64 {
	return 6.0 * c.sq.side * c.sq.side
}

// simple function.
func info(x shape) {
	fmt.Printf("\n%T \t%d\n", x, x)
	fmt.Println("area:", x.area())
}

func main() {

	// initialize shapes
	sq1 := square{10}
	sq2 := New(42) // using a constructor-like function
	cube1 := cube{sq1}
	ci := circle{10}

	// print shapes
	info(&sq1) // generate's pointer to the operand. return &{10} shows struct is a pointer
	info(&sq2)
	info(ci)
	info(&cube1)
}
