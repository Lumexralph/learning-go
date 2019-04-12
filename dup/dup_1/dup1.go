// dup1 prints the text of each line that appears more
// than once in the standard input, preceded by its count

package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
	counts := make(map[string]int)

	// Scanner that reads input and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)

	// we are not handling error at the moment
	// reads the next line and removes the new line character from the end
	for input.Scan() {
		// retrieve the text
		line := input.Text()
		// gets the line, use it as the key and because
		// the value would be empty when the key was added,
		// the value will default to its zero value, which for
		// this  work is an int which will be 0 and then it is
		// incremented by 1
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}