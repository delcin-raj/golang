package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
		"Usage: %s network-type service \n",
		os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	service := os.Args[2]
	port,err := net.LookupPort(networkType,service)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error ",err.Error())
		os.Exit(1)
	}
	fmt.Println("service port ",port)
	os.Exit(0)
}

