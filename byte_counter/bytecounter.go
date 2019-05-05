package main

import "fmt"

type ByteCounter int

// add a method to the customized type
func (c *ByteCounter) Write(p []byte) (int, error) {
	// because you can't add 2 values of different types,
	// you have to convert int to ByteCounter type
	*c += ByteCounter(len(p))

	return len(p), nil
}

func main() {
	// having the method needed by the io.Writer interface
	// we can use the data type in any method or function that
	// needs the interface to be confident of executing its operation
	var c Lumex
	c.Write([]byte("Hello World!!"))

	fmt.Println(c)
}


type Lumex string

func (s *Lumex) Write(p []byte) (int, error)  {
	*s += (Lumex("mexy") + Lumex(p))

	return len(p), nil
}