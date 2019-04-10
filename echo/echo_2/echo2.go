package main

import (
	"fmt"
	"os"
)

func main() {
	// since it is the short variable declaration,
	// the type of the variable is inferred from the type
	// of the value
	s, sep := "", ""
	args := os.Args

	// using a range over the iterable, returns the index and value
	// because we don't need the index, we use a blank identifier _
	for _, arg := range args {
		s = sep + arg
		sep = " "
	}

	fmt.Println(s)
}