// server2 A minimal "echo" server

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {

	// create a route and the route handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	// log out to the stream or std out
	// behind the scenes, the server runs the handler (function) for each incoming request in a separate goroutine
	// so that it can server multiple requests simultaneously
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	// if two concurrent requests try to update count, at the same time it would lead to
	// inconsistencies, a bug called race conditions
	// to avoid this, we lock the current goroutine to access
	// the variable at least once at a time
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, " URL.Path  = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
