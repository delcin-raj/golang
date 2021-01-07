package main
import(
	"fmt"
	"io/ioutil"
	"net"
	"os"
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

