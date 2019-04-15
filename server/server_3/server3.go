// server2 A minimal "echo" server

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

func main() {

	// create a route and the route handler
	http.HandleFunc("/", handler)

	// log out to the stream or std out
	// behind the scenes, the server runs the handler (function) for each incoming request in a separate goroutine
	// so that it can server multiple requests simultaneously
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}

	fmt.Fprintf(w, "Host = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print("err")
	}
	for key, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}
