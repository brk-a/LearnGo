package main

import "fmt"

// primitive #3: select

func primitivesSelect()  {
	myChan := make(chan string)
	anotherChan := make(chan string)

	go func() {
		myChan <- "data 1"
	}()
	go func() {
		anotherChan <- "data 2"
	}()

	select {
	case msg:=<-myChan:
		fmt.Println(msg)
	case msg:=<-anotherChan:
		fmt.Println(msg)
	}
}