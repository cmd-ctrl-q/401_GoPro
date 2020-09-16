/* 
https://gobyexample.com/string-formatting

formatting string. 
- go offers several printing verbs to format values
-- eg. %d, %v, %o, %X, %p, %q, %s, %E, ...

...

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
	%o`, i,i,i,i,i,i)
}


