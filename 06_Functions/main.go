/*
	Functions, et al.

	Returning multiple values
	Variadic functions 
	-- accepts any number of arguments in a form of a slice or array
	Slice parameters

*/

package main 


// 




import "fmt"

func sum(si ...float64) (int, float64) {
	var length int
	var total float64

	for _, v := range si {
		length += 1
		total += float64(v)
	}
	return length, total
}

func main() {

	nums := []float64{1,2,3,4,5,6,7,8,9}
	length, total := sum(nums...)
	
	var average float64 = total / float64(length)
	fmt.Println(average)
}