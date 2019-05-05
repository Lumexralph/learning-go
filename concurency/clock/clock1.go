// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// the server is sequential, handling on client connection
	// or request at a time until we run every client connection
	// in different goroutines
	for {
		// Accept method blocks until an incoming request is made
		// then creates the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue // don't stop the program
		}

		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // maybe client is disconnected
		}
		time.Sleep(1 * time.Second)
	}

}
