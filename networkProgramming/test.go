package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter")
	reader := bufio.NewReader(os.Stdin)
	for {
		result,_ := reader.ReadBytes('\n')
		fmt.Println(string(result))
	}
}

