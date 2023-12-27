package main

import (
	"fmt"
)

func main()  {
	sendSoFar := 430
	const sendsToAdd = 25

	sendSoFar = incrementSends(sendSoFar, sendsToAdd)

	fmt.Println("you have sent %v messages", sendSoFar)
}

func incrementSends(sendSoFar, sendsToAdd int) int {
	sendSoFar += sendsToAdd
	
	return sendSoFar
}