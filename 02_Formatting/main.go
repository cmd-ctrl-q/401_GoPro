/*
https://gobyexample.com/string-formatting
https://golang.org/pkg/fmt/

Formatting strings
- go offers several printing verbs to format values
-- %d 	integer
-- %f 	float
-- %t 	boolean
-- %c 	char code
-- %x 	hex code
-- %b 	binary
-- %e	scientific notation
-- %s 	string
-- %v 	default value. Useful for float, int, and bool

and many others:
-- %o, %X, %p, %q, %s, %E, %v...
*/

package main

import "fmt"

func main() {

	i := 42
	fmt.Printf(`
	%d
	%T 
	%c
	%b
	%#X
	%o`, i, i, i, i, i, i)

	// Alternatively written
	fmt.Printf("%d\n%T\n%c\n%b\n%#X\n%o", i, i, i, i, i, i)

	str := "Hello"
	str1 := "World"

	fmt.Print("\n\n", str+str1, i, "\n")
	fmt.Println(i + i)

}
