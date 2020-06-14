package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var paramIP string
var paramPort int
var paramTimeout int
var paramUsername string

func init() {
	flag.StringVar(&paramIP, "ip", "127.0.0.1", "Please set ip address for connection, default is 127.0.0.1")
	flag.IntVar(&paramPort, "port", 25, "Default port is 25")
	flag.IntVar(&paramTimeout, "timeout", 500, "Timeout in ms")
	flag.StringVar(&paramUsername, "username", "user", "Default username is user")
}

func main() {
	flag.Parse()

	target := fmt.Sprintf("%s:%d", paramIP, paramPort)

	conn, err := net.DialTimeout("tcp", target, time.Duration(paramTimeout)*time.Millisecond)

	if err != nil {
		log.Fatal(err)
	}

	//If connection is fine, read and print banner
	banner, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(banner)

	//Try to verify user existence
	conn.Write([]byte("VRFY " + paramUsername + "\n"))
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
	conn.Close()
}
