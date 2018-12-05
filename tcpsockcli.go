package main

import (
        "fmt"
        "net"
        "os"
	"strconv"
)


func main() {
        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
                os.Exit(1)
        }
        service := os.Args[1]

	conn, err := net.Dial("tcp", service)
        checkError(err)
	defer conn.Close()

	for i := 1; i<=2; i++ {
        	bufSend  := make(net.Buffers, 10)
	        for a := range bufSend { bufSend[a] = []byte("loop" + strconv.Itoa(a)) }
		//num, err := conn.Write(bufSend) 
		num, err := bufSend.WriteTo(conn)	
		checkError(err)
		fmt.Fprintf(os.Stdout, "wrote %d bytes \n", num)
	        for a := range bufSend { bufSend[a] = []byte("loop" + strconv.Itoa(a)) }

	}
        os.Exit(0)
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error %v", err)
                os.Exit(1)
        }
}

