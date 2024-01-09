package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <- chan time.Time)  {
	for{
		select {
		case <- snapshotTicker:
			takeSnapshot()
		case <- saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot()  {
	fmt.Println("Taking backup snapshot... ")
}

func saveSnapshot() {
	fmt.Println("All back-ups saved")
}

func waitForData()  {
	fmt.Println("Awaiting data... ")
}

func test()  {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("=============================================================")
}

func main()  {
	test()
}