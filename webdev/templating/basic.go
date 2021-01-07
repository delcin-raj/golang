package main
import (
	"os"
	"log"
	"strings"
	"io"
)

func main() {
	name := os.Args[1]
	res := "<h1>"+name+"</h1>"
	nf,err := os.Create("name.html")
	if (err !=nil) {
		log.Fatal("Unable to create file")
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(res))
}

