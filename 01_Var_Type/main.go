/* 

## variables and types
- uint8, uint16, uint32, uint64, int8, int16, int32, int64, etc 
- float32, float64, rune, atom

*/

package main

import "fmt"

func main() {

	var

	// constant 
	const PI float64 = 3.14159265 

	// declare a var
	var number int // has initial value of 0
	fmt.Println(number)

	// initialize shorthand 
	batman := "not Bruce"
	fmt.Println(batman)

	// declare many
	var one, two, three int
	one = 1
	// one, two, three do not share addresses
	fmt.Println(one, two, three) // 1 0 0

	// declare many
	const (
		a = 1
		b = a + iota
		c
		_
		d
	)
	// iota represents the next successive integer starting at 1
	fmt.Println(a) // 1
	fmt.Println(b) // 1+1 = 2
	fmt.Println(c) // 3
	fmt.Println(d) // 5

	// initialize many
	const x, y, z int32 = 1, 2, 3

	// infer mixed types 
	const l, m, n = "lion", 42, false

}