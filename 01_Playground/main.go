// https://www.youtube.com/watch?v=CF9S4QZuV30

package main 

import "fmt"

func main() {


	const pi float64 = 3.14159265 

	// declare multiple diff variables 
	var (
		a = 2
		b = 3
	)
	fmt.Println(a+b)
	// strings: "double quotes" or `backtick `
	var name string = `batman` 
	fmt.Println(name)

	// 
	var boolean bool = true
	fmt.Printf("%.3f %T %t\n", pi, pi, boolean)

	// 
	fmt.Printf("digit %d\n", 100)
	fmt.Printf("byte %b\n", 100)
	fmt.Printf("char %c\n", 100)
	fmt.Printf("hex %x\n", 100)
	fmt.Printf("octal %o\n", 100)
	fmt.Printf("scientific notation %e\n", pi)


	// logical operators 
	// && || !

	// for loop
	i := 1 
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}

	// relational operators 
	// == != < > <= >= 

	// for loop
	for j := 0; j < 5; j++ {
		fmt.Printf("%d ", j)
	}

	// switch statement
	age := 5
	switch age {
		case 16 : fmt.Println("Go Drive")
		case 18 : fmt.Println("Go Vote")
		default : fmt.Println("Go have fun")
	}

	// arrays 
	var myArr[5] float64 
	myArr[0] = 163
	myArr[1] = 345345
	myArr[2] = 45
	myArr[3] = 3.1415
	myArr[4] = 1.618


	myArr2 := [5]float64 {1,2,3,4,5}
	fmt.Println(myArr2)

	for _, v := range myArr2 {
		fmt.Println(v)
	}

	// slice (like an array but no need to specify size)
	mySlice := []int{5,4,3,2,1}
	// slice of a slice 
	mySlice2 := mySlice[3:5]

	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(mySlice[:2])

	// dynamic sized array using make()
	// slice without a defined set of initial values 
	// make(type, default of zero for first 5, max size of 10)
	mySlice3 := make([]int, 5, 10)
	fmt.Println(mySlice3)
	// append
	mySlice3 = append(mySlice3, 0, -1) // append(value1,value2,value3,...)
	fmt.Println(mySlice3)


	// MAP
	mymap := make(map[string] int)
	mymap["TheodoreRoosevelt"] = 42 
	fmt.Println(mymap["TheodoreRoosevelt"])
	fmt.Println(len(mymap))
	mymap["John F. Kennedy"] = 43 
	fmt.Println(len(mymap))
	// delete key 
	delete(mymap, "John F. Kennedy")
	fmt.Println(len(mymap))


	// functions 
	num1, num2 := getNext2(10)
	fmt.Println(num1, num2)

	fmt.Println(subtractAll(1,2,3,4,5))

	// closure - function portability. 
	// assigning a function to a variable to retain it and its variables. 
	// A closure is a persistent scope which holds on to local variables 
	// even after the code execution has moved out of that block. 
	// The scope object and all its local variables are tied to the function 
	// and will persist as long as that function persists.
	num3 := 3 
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	fmt.Println(summation(10))
	fmt.Println(factorial(5))


	// DEFER 
	defer printTwo()  // printed last (ie very end of function)
	printOne()

	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	demPanic()

	// POINTERS 31:04 


}

func printOne() {fmt.Println(1)}
func printTwo() {fmt.Println(2)}

func getNext2(n int) (int, int) {
	return n+1, n+2
}

func subtractAll(args ...int) int {
	finalValue := 0 
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

func summation(num int) int {
	if num == 1 {
		return num
	} 
	return (num + summation(num - 1))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} 
	return num * factorial(num - 1)

}

func safeDiv(num1, num2 int) int {
	defer func() {
		// recover catch error if one occurs but does not crash program.
		// ie continues to run program
		fmt.Println(recover()) // Println is not needed
	}() 
	// if num2 is a zero, it will run the defer func before returning solution.
	// but since num2 = 0, recover will run 
	solution := num1 / num2 
	return solution 
}

func demPanic() {

	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC!")
}