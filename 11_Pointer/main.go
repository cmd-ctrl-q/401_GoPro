/* 
	everything in go is Pass By Value, 
	even when you pass a memory address, you are passing a value
	pointers allow you to move large data by passing around an address
*/

package main 

import "fmt"

func main() {

	a := 42 

	// referencing 
	var b = &a // store pointer in b 
	fmt.Printf("\n%T\t%p\n\n", b,b) 	// print pointer
	
	// dereference 
	fmt.Println(*b)
	*b = 22  // change value of `b` to 22 
	fmt.Println(a) 

}