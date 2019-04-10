package main

import (
	"fmt"
	"os"
)


func main() {
	// if an empty value is created, it is
	// implicitly initialized to the zero value of the type
	var s, sep string

	args := os.Args

	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
}