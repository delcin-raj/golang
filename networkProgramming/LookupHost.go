package main

import (
	"fmt"
	"net"
)

func main() {
	addrs,_ := net.LookupHost("www.instagram.com")
	for _,s := range addrs {
		fmt.Println(s)
	}
}
