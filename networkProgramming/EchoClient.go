package main
import(
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
	service := ":1200"
	tcpAddr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)

	// Client should dial 
	// while the server listens
	conn, err := net.DialTCP("tcp",nil,tcpAddr)
	// Dial prompts the server to accept
	// Our server will automatically provide it's only service
	checkError(err)
	reader := bufio.NewReader(os.Stdin)
	var buf [512]byte
	for {
		fmt.Println("Enter text")
		text,_ := reader.ReadBytes('\n')
		_,err := conn.Write(text)
		checkError(err)
		n,err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println("Echoed from server: ",string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
                os.Exit(1)
        }
}

