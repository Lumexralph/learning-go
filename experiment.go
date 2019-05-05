package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello world from main function")

	var u uint8 = 255

	fmt.Println(u, u+1, u*u)

	var i int8 = 127

	println(i, i + 1, i * i)

	var x uint8 = 1 << 1 | 1 << 5 | 1 << 6
	var y uint8 = 1 << 1 | 1 << 2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n",  x & y) // intersection
	fmt.Printf("%08b\n",  x | y) // union of both
	fmt.Printf("%08b\n",  x ^ y) // symmetric difference
	fmt.Printf("%08b\n",  x &^ y) // symmetric difference

	for i := uint(0); i < 8; i++ {
		
	}

}

func init()  {
	fmt.Println("Hello world from init function")
}