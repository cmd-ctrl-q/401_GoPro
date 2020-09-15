/* 

	Go uses codepoints to represent strings and chars.
	Go assumes strings are a UTF8 encoded series of bytes

*/

package main 

import "fmt"

func main() {
	a := "Äpfel"

	// Codepoint.
	// uses 2 codepoints to represent Ä
	fmt.Println([]byte(a)) 

	// uses 1 codepoint to represent Ä
	fmt.Println([]rune(a))

	// UTF-8 (hexcode)
	fmt.Printf("% x\n", a) 

	// Glyph (returns string)
	fmt.Printf("%q \n", a) 

	// Unicode codepoint
	fmt.Printf("%+q \n", a) 
	fmt.Printf("%U \n", []rune(a))

}