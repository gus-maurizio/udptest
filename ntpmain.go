package main

import (
	"fmt"
	"github.com/gus-maurizio/ntp"
	"time"
)

func main() {
	for {
		fmt.Printf("Looping NTP checks at %v \n", time.Now())
		options := ntp.QueryOptions{Timeout: 300 * time.Millisecond}
		timegoo, errgoo := ntp.QueryWithOptions("time.google.com", options)
		timenis, errnis := ntp.QueryWithOptions("time.nist.gov", options)

		// timegoo, errgoo := ntp.Time("time.google.com")
		fmt.Printf(" goo %+v err: %v\n", timegoo, errgoo)
		fmt.Printf(" nis %+v err: %v\n", timenis, errnis)
		fmt.Printf("Waiting 5s at %v \n", time.Now())
		time.Sleep(5 * time.Second)
	}
}
