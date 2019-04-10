package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	// to avoid the situation when the value of
	// data or the slice might increase, to make
	// an improvement on the created echo implementation
	args := os.Args[1:]
	s := strings.Join(args, " ")

	fmt.Println(s)
}