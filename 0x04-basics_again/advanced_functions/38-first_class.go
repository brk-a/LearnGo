package main

import (
	"fmt"
)

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}

	for _, message:=range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}

	return formattedMessages
}

func addSignature(message string) string {
	return message + "Thank you."
}

func addGreeting(message string) string {
	return "Hi." + message
}

func test(messages []string, formatter func(string) string)  {
	defer fmt.Println("===========================================================")

	formattedMessages := getFormattedMessages(messages, formatter)
	if len(formattedMessages)!=len(messages) {
		fmt.Println("the number of messages returned is incorrect.")
		return
	}

	for i, message:=range messages {
		formatted := formattedMessages[i]
		fmt.Printf("* %s -> %s\n", message, formatted)
	}
}

func main() {
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend, however, I am busy.",
		"Got any hot stock tips?",
	}, addSignature)
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend, however, I am busy.",
		"Got any hot stock tips?",
	}, addGreeting)
}