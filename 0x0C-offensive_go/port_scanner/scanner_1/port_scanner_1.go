package port_scanner_1

import (
	"fmt"
	"net"
)

func PortScanner_1() {
	for i := 0; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue //port may be closed or otherwise unavailable
		}
		conn.Close()
		fmt.Printf("Port open: %d\n", i)
	}
}
