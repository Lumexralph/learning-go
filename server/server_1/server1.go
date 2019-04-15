// A minimal "echo" server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create a route and the route handler
	http.HandleFunc("/", handler)

	// log out to the stream or std out
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " URL.Path  = %q\n", r.URL.Path)
}
