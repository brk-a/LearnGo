package main

import "fmt"

type messageToSend struct {
	phoneNumber string
	message string
}

func testStruct(m messageToSend)  {
	fmt.Println("Sending message %v to %v", m.message, m.phoneNumber)
	fmt.Println("==================================================================")
}

func main() {
	testStruct(messageToSend{
		phoneNumber: "0712345678",
		message: "You have signed up for Textio",
	})
	testStruct(messageToSend{
		phoneNumber: "0701234567",
		message: "You have subscribed to the premium package on Textio",
	})
	testStruct(messageToSend{
		phoneNumber: "0723456789",
		message: "You have logged out of Textio",
	})
}