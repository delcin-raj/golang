package main
import  (
	"net"
	"fmt"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d",p)
		conn,err := net.Dial("tcp",address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int,100)
	results := make(chan int) // send the result to main routine
	var openports []int 
	num_of_workers := 68
	for i := 0; i < num_of_workers; i++ {
		go worker(ports,results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	
	// result gathering part
	for i := 0; i< 1024; i++ {
		port := <- results
		if port != 0 {
			openports = append(openports,port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openports)
	for _,port := range openports {
		fmt.Println(port,": is open")
	}
}
