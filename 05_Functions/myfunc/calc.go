/*
	*** Functions ***

	First Class Citizens / Higher Order Functions
	- The ability to treat functions as regular values
	- e.g.
		- assign functions to varibles
		- pass functions as arguments
		- return functions

	Other:
	Functions can return multiple variables
	Functions can be stored as struct fields.

	Variadic functions:
	- accepts any number of arguments in a form of a slice or array
	- denoted as: func(slice ...type)
*/

package myfunc

import (
	"math/rand"
)

// PickRandom picks two random numbers in a slice of ints
func PickRandom(min, max int) (int, int) {
	rn1 := rand.Intn(max-min) + min
	rn2 := rand.Intn(max-min) + min
	return rn1, rn2 // return multiple values
}

type fn func(int, int) (int, int) // create type func

// finds difference between two values.
func CalcDiff(f fn, min, max int) (int, int, int) {
	// pick random ints between min and max values
	rn1, rn2 := f(min, max)

	// find difference and return (rn1, rn2, difference)
	if rn1 > rn2 {
		return rn2, rn1, (rn1 - rn2)
	}
	return rn1, rn2, (rn2 - rn1)
}
