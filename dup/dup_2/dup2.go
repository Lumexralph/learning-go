// dup2 prints the text of each line that appears more
// than once in the standard input.
// It reads from stdin or from a list of named files
// dup2 operates in a ‘‘streaming’’ mode in which input is read
// and broken into lines as needed, so in principle it can handle
// an arbitrary amount of input, it does not read everything into
// memory

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// go through each file
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				// output the error and continue to the next file
				fmt.Fprintf(os.Stderr, "dup_2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// close the file to free up space from the memory and
			// avoid memory leakage
			f.Close()
		}
	}

	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
