package port_scanner_0

import (
	"fmt"
	"net"
)

func PortScanner_0() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Printf("Port open: %d\n", 80)
	}
}
