package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

// receive from channel and print data
func receive(e, o, quit <-chan int) {
	for {
		select {
		// if receiving from channel e,
		// then store value in v
		// and print it
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		// if receiving from channel o,
		// then store value in v
		// and print it
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		// if receiving from channel q,
		// then store value in v
		// and print it
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			} else {
				fmt.Println("from comma ok:", i)
			}
		}
	}
}

// send/add data to channel
func send(e, o, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(quit)
}
