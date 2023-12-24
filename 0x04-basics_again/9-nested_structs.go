package main

type message struct {
	message string
	sender user
	recepient user
}

type user struct {
	name string
	number string
}

func canSendMessage(mToSend messageToSend) bool {
	
}