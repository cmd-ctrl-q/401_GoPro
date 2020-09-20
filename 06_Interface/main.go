/*

	*** Interfaces ***

	Struct embedding
		Embedding a struct in another struct to expose the embedded structs
		fields and methods for a struct to use.
	Mixins
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

// square is promoting its fields to threeDSquare
type threeDSquare struct {
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

// Mixins
// implements shape by using the signature area() from shape
func (s *square) area() float64 { // takes in pointer
	return s.side * s.side
}

// implements shape by using area() method from shape
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 3dsquare implements shape
func (tds threeDSquare) area() float64 {
	return 6.0 * tds.sq.side * tds.sq.side
}

// simple function.
func info(x shape) {
	fmt.Println(x)
	fmt.Println(x.area())
}

func main() {

	// initialize shapes
	sq := square{10}
	sq2 := New(42) // using a constructor-like function
	threeDsq := threeDSquare{sq}
	ci := circle{10}

	// print shapes
	info(&sq) // generate's pointer to the operand. return &{10} shows struct is a pointer
	info(&sq2)
	info(ci)
	info(&threeDsq)
}
