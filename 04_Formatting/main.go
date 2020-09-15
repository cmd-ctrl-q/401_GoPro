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