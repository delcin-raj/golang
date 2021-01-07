package main
import (
	"fmt"
	"log"
	"os"
	"io"
)

// defines a type
type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int,error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}
// declaring that the type Fooreader struct has a method Read

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int,error) {
	fmt.Print("out>")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)
	_,err := io.Copy(&writer,&reader)
	if err != nil {
		log.Fatalln("unable to read/write data")
	}
}

