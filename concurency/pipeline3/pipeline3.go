package main

import (
	"fmt"
	"time"
)

// the sending goroutine
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// have a channel that will send data and also a
// chan that will receive from the chan out
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}

	close(out)
}

// printer receives a chan that will receive value
func printer(in <-chan int) {
	for v := range in {
		fmt.Println("The received value", v)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	// create the channels
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}
