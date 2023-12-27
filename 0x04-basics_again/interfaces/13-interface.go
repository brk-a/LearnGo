package main

import (
	"fmt"
	"time"
)

func sendMessage(msg messageSent)  {
	fmt.Println(msg.getMessage())
}

type messageSent interface {
	getMessage() string
}

type specialMessage struct {
	specialTime time.Time
	recipientName string
}

func (sm specialMessage) getMessage() string {
	return fmt.Sprintf("Hi, %s. It's %s. Congratulations!", sm.recipientName, sm.specialTime)
}

type sendingReport struct {
	reportName string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You have now sent %v messages`, sr.reportName, sr.numberOfSends)
}

func testIn(m messageSent) {
	sendMessage(m)
	fmt.Println("==========================================================")
}

func main() {
	testIn(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	testIn(specialMessage{
		recipientName: "Goat Matata",
		specialTime:  time.Date(1970, 01, 11, 0, 0, 0, 0, time.UTC),
	})
	testIn(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	testIn(specialMessage{
		recipientName: "Mzee Kobe",
		specialTime:  time.Date(1900, 01, 01, 0, 0, 0, 0, time.UTC),
	})
}