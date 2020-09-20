package myfunc

import "fmt"

// ExamplePickRandom picks two random values between 0 and 10
func ExamplePickRandom() {
	PickRandom(0, 10)
	// Output
	// 3
}

// ExampleCalcDiff uses some function fn to find the difference between two numbers.
func ExampleCalcDiff() {
	fmt.Println(CalcDiff(PickRandom, 0, 10))
	// Output:
	// min: 		1
	// max: 		7
	// difference: 	6
}
