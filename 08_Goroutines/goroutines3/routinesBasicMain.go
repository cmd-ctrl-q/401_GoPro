/*
	Goroutines

	Example of using multiple goroutines
*/

package main

import (
	"fmt"
	"time"
)

func hello(alert chan string) {
	for i := 0; ; i++ {
		alert <- "Hello? You there? Press Enter if you're there!"
	}
}

func print(alert chan string) {
	for {
		text := <-alert
		fmt.Println(text)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var alert chan string = make(chan string)

	go hello(alert)
	go print(alert)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Oh hi there! Good bye!")
}
