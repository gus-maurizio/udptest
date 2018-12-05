package main

import (
	"io"
	"log"
	"net"
	"time"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	l, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		log.Printf("%d connected from %#v\n", time.Now().UnixNano(), conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			numc, _ := io.Copy(c, c)
			log.Printf("%d read %d  from %#v\n", time.Now().UnixNano(), numc, c.RemoteAddr())
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
