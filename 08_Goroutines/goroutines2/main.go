/*

	Goroutines

	Example of using multiple goroutines

	There is always one main goroutine for a program.
	Using `go` you can add many more.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("cpu", runtime.NumCPU())               // cpu count
	fmt.Println("go routines", runtime.NumGoroutine()) // existing goroutines

	// WaitGroup waits for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2) // adds two more routines

	go func() {
		fmt.Println("func 1 called")
		wg.Done()
	}()

	go func() {
		fmt.Println("func 2 called")
		wg.Done()
	}()

	fmt.Println("\ncpu", runtime.NumCPU())             // cpu count
	fmt.Println("go routines", runtime.NumGoroutine()) // existing goroutines

	wg.Wait() // unlocks both goroutine

	fmt.Println("\ncpu", runtime.NumCPU())             // cpu count
	fmt.Println("go routines", runtime.NumGoroutine()) // existing goroutines
}
