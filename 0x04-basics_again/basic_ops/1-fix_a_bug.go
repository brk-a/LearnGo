package main

import (
	"fmt"
)

func textMessages()  {
	messagesFromAlice := []string{
		"Doing anything later?",
		"Did you get my last message?",
		"Don't leave me hanging",
		"Please respond",
	}

	numMessages := float64(len(messagesFromAlice))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages

	fmt.Println("Total cost is %.2f\n", totalCost)
}

func main()  {
	textMessages()
}