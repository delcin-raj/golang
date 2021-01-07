package main
import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) !=2 {
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	// her we expect domain name for the ip
	name := os.Args[1]
	addr,err := net.ResolveIPAddr("ip6",name)
	if err!= nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ",addr.String())
	os.Exit(0)
}



