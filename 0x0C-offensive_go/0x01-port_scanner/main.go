package main

import (
	"fmt"

	scanner_0 "github.com/brk-a/offensive_go/port_scanner/scanner_0"
	scanner_1 "github.com/brk-a/offensive_go/port_scanner/scanner_1"
	scanner_2 "github.com/brk-a/offensive_go/port_scanner/scanner_2"
)

func main(){
	fmt.Printf("Running single-port scanner...\n")
	scanner_0.PortScanner_0()
	fmt.Printf("=================================================\n")
	fmt.Printf("Running port range scanner sans concurrency...\n")
	scanner_1.PortScanner_1()
	fmt.Printf("=================================================\n")
	fmt.Printf("Running port range scanner with concurrency...\n")
	scanner_2.PortScanner_2()

}