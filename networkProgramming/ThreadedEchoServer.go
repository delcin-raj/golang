package main

import (
	"fmt"
	"net"
	"os"
//	"io/ioutil"
)

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
                os.Exit(1)
        }
}

func main() {
	service := ":1200"
	tcpAddr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)

	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)

	for {
		// Try to create a connection
		conn,err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// We have to wait for the the client to write the command
	defer conn.Close()
	var buf [512]byte
	for {
		n,err := conn.Read(buf[0:])
		if err != nil {
			// terminate the connection
			return
		}
		fmt.Println(string(buf[0:]))
		_,err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
