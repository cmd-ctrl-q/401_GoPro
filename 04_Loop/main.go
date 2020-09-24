/*
	Looping
*/
package main

import (
	"fmt"
)

func main() {

	// regular for loop
	fmt.Println("even\todd")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\t", i)
		} else {
			fmt.Printf("%d\n", i)
		}
	}

	// while loop
	var i, b = 0, true
	for b {
		if i == 5 {
			b = false
		}
		fmt.Println(b)
		i += 1
	}

	// range over a slice
	// slice is like array but not necessary to define size
	arr := []float64{1, 2, 3, 4, 5} // slice
	for _, v := range arr {
		fmt.Printf("$%.2f\t", v)
	}
}
