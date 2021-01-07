package main
import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// All network interfaces in the system will listen via
	// port 1200
	service := ":1200"
	/* type TCPAddr struct {
		ip IP
		port int
		Zone string
	}
		
	*/
	tcpAddr, err := net.ResolveTCPAddr("tcp",service)
	checkError(err)

	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)

	for {
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
                os.Exit(1)
        }
}

