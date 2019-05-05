package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		// Counter
		for x := 0; x < 150; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}

		close(squares)
	}()

	// Printer in the main goroutine
	for x := range squares {
		fmt.Println("The final pipeline value ", x)
	}
}
