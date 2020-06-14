package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var paramIP string
var paramPort int
var paramTimeout int

func init() {
	flag.StringVar(&paramIP, "ip", "127.0.0.1", "Please set ip address for connection, default is 127.0.0.1")
	flag.IntVar(&paramPort, "port", 25, "Default port is 25")
	flag.IntVar(&paramTimeout, "timeout", 500, "Timeout in ms")
}

func main() {
	flag.Parse()

	target := fmt.Sprintf("%s:%d", paramIP, paramPort)

	conn, err := net.DialTimeout("tcp", target, time.Duration(paramTimeout)*time.Millisecond)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port is open, lets try to send VRFY command")
	conn.Close()
}