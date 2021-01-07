package main
import (
	"fmt"
	//"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// wg is controlled by the main thread
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			//the corputine is deferring the call to Done
			//until it's work was done
			address := fmt.Sprintf("scanme.nmap.org:%d",i)
			conn,err := net.Dial("tcp",address)
			if err == nil {
				fmt.Println("port:",i," is open")
				conn.Close()
			}
			fmt.Println(j)
		}(i)
	}
	wg.Wait()
}

