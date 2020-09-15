package main 

import "fmt"


func main() {

	// 
	fmt.Println("even\todd")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\t",i)
		} else {
			fmt.Printf("%d\n",i)
		}
	}

	// 
	var b bool = true
	for b {
		fmt.Println()
	}

	//
}