/*
	ways to limit the scope of variables is to use closures
*/

package main

import "fmt" 

func main() {

	x := 42

	// simple closure. 
	{
		fmt.Println(x) 
		y := "words"
		fmt.Println(y) 
		// defer prints last
	}

	// closure with anon func
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment()) // 43
}
