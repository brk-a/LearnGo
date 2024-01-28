package main

import "fmt"

// primitive #2: channels

func primitivesChannels()  {
	myChan := make(chan string)

	go func() {
		myChan <- "data 1"
	}()

	msg :=  <- myChan

	fmt.Println(msg)
}