package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "shit", "s***")
	messageVal = strings.ReplaceAll(messageVal, "fuck", "f***")

	*message = messageVal
}

func test(messages []string) {
	for _, message:=range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well, this is shit!",
		"damn robots",
		"fuck them to hell",
	}
	messages2 := []string{
		"I fucking swear all the fucking time",
		"proclivity to employ profanity is a sign of a weak fucking mind",
		"that horseshit is not fucking real",
	}

	test(messages1)
	test(messages2)
}