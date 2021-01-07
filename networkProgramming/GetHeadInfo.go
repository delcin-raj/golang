package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage: %s host:port ",os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	// resolving the ip address of the TCPAddr struct
	tcpAddr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)

	// Establishing a dupplex connection between the service and
	// the program
	conn, err := net.DialTCP("tcp",nil,tcpAddr)
	checkError(err)

	// Writing the request to the server
	_,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// Reads everything until the end of file
	result,err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}
