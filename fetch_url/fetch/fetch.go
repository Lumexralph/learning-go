// Fetch prints the content found at a URL.

package main

import (
	"fmt"
	"io"
	"strings"

	// "io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// check if the url has a prefix
		result := strings.HasPrefix(url, "http://")
		if !result {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// displays the error and close the current process
			os.Exit(1)
		}
		// read the entire the body stream
		// byte, err := ioutil.ReadAll(resp.Body)

		// to avoid ensuring a buffer large to take // the entire stream of response
		fmt.Printf("The returned status of reponse %d", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)

		// close the body after reading the value to free resources and avoid memory leakage
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)

			// end the current process, to indicate the process exited with error
			// when the process exits with only 0, it means there was not error
			os.Exit(1)
		}
		// write a response to the standard output
		// fmt.Printf("%s", byte)
	}
}
