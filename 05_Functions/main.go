/*
	Functions 

	First class citizens
	- Then assigned to variables and passed as arguments to other functions and can be returned from a function. 
	- Functions can return multiple variables	
	
	
	Functions can be stored as struct fields.

	Variadic functions 
	-- accepts any number of arguments in a form of a slice or array
	Slice parameters
	Anon func
*/

package main 

import (
	"fmt"
)


func print1(f func()) {
}

// func print2(f func()) {
// 	f()
// }

func main() {


}




// the ...type notation unfurls container types (eg slice, array, maps)
// func sum(si ...float64) (int, float64) {
// 	var length int
// 	var total float64

// 	// Blank Identifier notation _
// 	// used for readability, and 
// 	for _, v := range si {
// 		length += 1
// 		total += float64(v)
// 	}
// 	return length, total
// }


	// nums := []float64{1,2,3,4,5,6,7,8,9}
	// length, total := sum(nums...)
	
	// var average float64 = total / float64(length)
	// fmt.Println(average)

	// // anon func
	// var x int = 42
	// increment := func() int {
	// 	x++
	// 	return x
	// }
