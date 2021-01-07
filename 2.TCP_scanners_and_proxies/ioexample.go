package main
import (
	"fmt"
	"log"
	"os"
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
	input := make([]byte,4096)

	s,err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data")
	}
	fmt.Println ("Read ",s, " bytes")

	s,err = writer.Write(input)
	if err != nil {
		log.Fatalln("write failed")
	}
	fmt.Println("Written ", s, " bytes")
}
		

