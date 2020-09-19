package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	///// 0 division
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	///// File I/O
	file, err := os.Create("filename.txt") // err is nil if there is no error

	if err != nil {
		fmt.Println("Could not open file")
		log.Fatal(err)
	}

	file.WriteString("This is some random text")
	file.Close()

	stream, err := ioutil.ReadFile("filename.txt")

	if err != nil {
		err = errors.New("You entered a bad name")
		fmt.Println(err)
		// log.Fatal(err) // "errors" do not automatically end a program, if you suspect errors to arise, utilize flow control
	}

	readString := string(stream)
	fmt.Println(readString)

	println("All lines executed.")
}

// Demo refer, panic, recover (use flow control for compile errors)
func safeDiv(num1, num2 int) int {
	defer func() { // anon function
		if r := recover(); r != nil {
			fmt.Println("Recovered in safeDiv(); ", r)
		}

	}()

	solution := num1 / num2
	return solution
}
