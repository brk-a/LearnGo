package main

import (
	"fmt"
	"time"
)

func doneChannel()  {
	done := make(chan bool)
	go run(done)
	time.Sleep(time.Second * 5)
	close(done)
	
}

func run(done <-chan bool) {
	for{
		select{
		case <- done:
			return
		default:
			fmt.Println("running...")
		}
	}
}