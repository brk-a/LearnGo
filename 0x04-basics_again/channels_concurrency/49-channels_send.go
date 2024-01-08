package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDbs int, dbChan chan struct{})  {
	for i:=0; i<numDbs; i++ {
		<- dbChan
	}
}

func test(numDbs int)  {
	dbChan := getDatabasesChannel(numDbs)
	fmt.Printf("Waiting for %v databases...\n", numDbs)
	waitForDbs(numDbs, dbChan)
	time.Sleep(time.Millisecond * 10)
	fmt.Println("All databases are online")
	fmt.Println("============================================")
}

func main()  {
	test(3)
	test(0)
	test(-1)	
}

func getDatabasesChannel(numDbs int) chan struct{} {
	ch := make(chan struct{})
	go func()  {
		for i:=0; i<numDbs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()

	return ch
}