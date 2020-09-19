/*
	everything in go is Pass By Value,
	even when you pass a memory address, you are passing a value
	pointers allow you to move large data by passing around an address
*/

package main

import "fmt"

func main() {

	x := 0

	changeXVal(x) // passes value of x, not x itself
	fmt.Println("x = ", x)

	changeXValNOW(&x) // pass in pointer address
	fmt.Println("Memory address for x = ", &x)
	fmt.Println("x = ", x)

}

func changeXVal(x int) {
	x = 2
}

func changeXValNOW(x *int) { // we are going to be sent a reference to the value and not the value itself, allows us
	// to change the value inside the memory address referenced by the pointer

	// store 2 in the memory address that x refers to
	*x = 2 // *x, has nothing to do with the address, but the value itself

}
