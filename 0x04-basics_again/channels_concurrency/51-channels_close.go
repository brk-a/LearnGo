package main

import (
	"fmt"
	"time"
)

func countReports(numSentCh chan int) int {
	count := 0
	for {
		numSent, ok := <- numSentCh
		if !ok {
			break
		}
		count += numSent
	}

	return count
}

func test(numBatches int)  {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Starting count...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent.\n", numReports)
	fmt.Println("===============================================")
}

func main() {
	test(3)
	test(0)
	test(-2)
	test(6)
}

func sendReports(numBatches int, ch chan int)  {
	for i:=0; i<numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond*100)
	}
	close(ch)
}