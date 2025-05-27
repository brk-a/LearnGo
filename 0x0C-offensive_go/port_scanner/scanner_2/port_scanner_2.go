package port_scanner_2

import (
	"fmt"
	"net"
	"sync"
)

func PortScanner_2() {
	var wg sync.WaitGroup
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			_, err := net.Dial("tcp", address)
			if err != nil {
				 //port may be closed or otherwise unavailable
			}
			// conn.Close()
			fmt.Printf("Port open: %d\n", j)
		}(i)
	}
	wg.Wait()
}
