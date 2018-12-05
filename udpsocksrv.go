package main

import (
        "fmt"
        "net"
        "os"
	"time"
)


func main() {
        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
                os.Exit(1)
        }
        service := os.Args[1]

	pc, err := net.ListenPacket("udp", service)
        checkError(err)
	defer pc.Close()
	// err = pc.SetReadDeadline(time.Time(1000000000))
        checkError(err)

        bufRead  := make([]byte, 32000)
	bufMsgs := make(net.Buffers,100)
	for a := range bufMsgs { bufMsgs[a] = make([]byte,300)}

	for {
		numbytes, sendaddr, err := pc.ReadFrom(bufRead) 
		checkError(err)
		// nummsgs,  merr := bufMsgs.Read(bufRead)	
		// checkError(merr)
		fmt.Fprintf(os.Stdout, "%d read %d bytes [%+v] in msgs from %+v\n", time.Now().UnixNano(), numbytes, bufRead[:numbytes], sendaddr)
	}
        os.Exit(0)
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error %v", err)
                os.Exit(1)
        }
}

