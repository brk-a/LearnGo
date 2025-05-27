package advanced_port_scanner

import (
	advanced_scanner "github.com/brk-a/offensive_go/advanced_port_scanner/advanced_port_scanner"
)


func Main(){
	advanced_scanner.AdvancedPortScanner(1000, 1, 65336, "localhost")
}