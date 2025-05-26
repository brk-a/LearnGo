package main

import (
	port_scanner "github.com/brk-a/offensive_go/port_scanner"
	remote_shell "github.com/brk-a/offensive_go/remote_shell"
)

func main(){
	port_scanner.Main()
	remote_shell.Main()
}