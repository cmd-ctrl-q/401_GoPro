/*Every Go executable starts with a package declaration which provides a
way to reuse code. The package name does not have to be the same as the
file name.*/
package main

// Import libraries and packages
// fmt is the formatting library
import (
	"fmt"
	"sync"
)

// The entry point of any go file is the main function
func main() {
	var x int      // default declaration, var name type
	var y int = 12 // explicit declaration + initialization
	z := 12        // recommended syntax, name = expression

	x, y, z = 10, 20, 30 // multi-variable declaration
	x, y = y, x          // swap the content of x and y

	// Printf and Println are functions inside the fmt library
	fmt.Println("x: ", x, "\t\ty: ", y, "\t\tz: ", z)
	fmt.Printf("x: %d\t\ty: %d\t\tz: %d\n", x, y, z)

	var wg sync.WaitGroup
	wg.Wait()
}
