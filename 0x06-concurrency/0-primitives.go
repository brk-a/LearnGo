package main

// primitive #1: go routines

import (
	"fmt"
	"time"
)

func primitivesGoRoutines()  {
	go printNum(1)
	go printNum(2)
	go printNum(3)

	time.Sleep(time.Second * 2)

	fmt.Println("this prints last")
}

func printNum(num int64) {
	fmt.Println(num)
}