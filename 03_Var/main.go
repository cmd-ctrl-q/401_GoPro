package main

import "fmt"

func main() {

	// constant 
	const pi float64 = 3.14159265 

	// initialize a var
	var number int // has initial value of 0
	fmt.Println(number)

	number = 42 // declare existing var
	fmt.Println(number)

	// short declaration 
	batman := "not Bruce"
	fmt.Println(batman)

	// multiple declarations 
	const (
		a = 1
		b = a + iota // represents the next successive integer
		c  // c = iota = 3
		_  // _ = iota = 4
		d  // d = iota = 5
	)
	fmt.Println(a) // 1
	fmt.Println(b) // 1+1 = 2
	fmt.Println(c) // 3
	fmt.Println(d) // 5
}