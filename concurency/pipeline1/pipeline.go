package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			// start the counting and pass
			// it to the naturals channel
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			// receive the value from the naturals channel
			x, ok := <- naturals
			if !ok {
				break // means the channel is drained and most likely closed
			}

			// send the square of x to squares channel
			squares <- x * x
		}

		close(squares)
	}()

	// Printer in the main goroutine
	for {
		fmt.Printf("The value from the pipeline %d\n", <-squares)

		time.Sleep(2 * time.Second)
	}
}
