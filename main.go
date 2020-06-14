package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var paramIP string
var paramPort int
var paramFile string

func init() {
	flag.StringVar(&paramIP, "ip", "127.0.0.1", "Please set ip address for connection, default is 127.0.0.1")
	flag.IntVar(&paramPort, "port", 25, "Default port is 25")
	flag.StringVar(&paramFile, "file", "file", "Filename with usernames")
}

func main() {
	flag.Parse()

	target := fmt.Sprintf("%s:%d", paramIP, paramPort)

	conn, err := net.Dial("tcp", target)

	if err != nil {
		log.Fatal(err)
	}

	//If connection is fine, read and print banner
	banner, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(banner)

	file, err := os.Open(paramFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//Try to verify user existence
		checkUsername(scanner.Text(), conn)
	}
	conn.Close()
}

// checkUsername
// Send VRFY username command to server for username checking
func checkUsername(name string, conn net.Conn) {
	conn.Write([]byte("VRFY " + name + "\n"))
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
