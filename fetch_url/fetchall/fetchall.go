// Fetchall fetches URLs in parallel and reports their
// times and sizes.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	// get current time
	start := time.Now()
	// create a channel of strings, that expects string
	ch := make(chan string)

	// go through all the URL
	for _, url := range args {
		// using go concurrency to send
		// multiple requests to different URL at
		// same time
		go fetch(url, ch) // start a goroutine
	}

	// receive or get data from the channel
	for range args {
		// receive data from the channel and send it
		// to the std out using the println method
		// which sends data to the std out
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

// fetch expects a url that is a string and a channel
// the channel expects to receive a string
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send the error to the channel
		return
	}

	// read the body of the response and discard it // by writing to the output stream of ioutil.Discard
	nbytes, err :=  io.Copy(ioutil.Discard, resp.Body)
	// close the body or the stream to free resources and avoid
	// memory leakage
	if err != nil {
		// send the error to the channel
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	// return the summary of the result to the main function goroutine
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)

}
