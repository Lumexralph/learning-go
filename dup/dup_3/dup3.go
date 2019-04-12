// dup3 reads the entire input into memory in one big gulp,
// split it into lines all at once, then process the line

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// NB: Under the hood, bufio.Scanner, ioutil.ReadFile and ioutil.WriteFile use
// the Read and Write Method of *os.File


func main() {
	counts := make(map[string]int)
	args := os.Args[1:]

	for _, filename := range args {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			// output the error into the stream of the current process
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// ReadFile returns a byte slice that must be converted into a
		// string so it can be split by strings.Split
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
