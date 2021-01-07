package main
import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, " Usage %s dottend-ip-addr ones bits\n,os.Args[0]")
		os.Exit(1)
	}
	// The net package deals with IPv4 and IPv6 seamlessly
	// Because CIDR works for both IPv4 and IPv6
	dotAddr := os.Args[1]
	ones,_ := strconv.Atoi(os.Args[2])
	bits,_ := strconv.Atoi(os.Args[3])

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := net.CIDRMask(ones,bits)
	fmt.Println(mask)
	network := addr.Mask(mask)
	fmt.Println("Address is ", addr.String(),
		"\n Mask length is ", bits,
		"\n Leading ones count is ",ones,
		"\n Mask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	os.Exit(0)
}
