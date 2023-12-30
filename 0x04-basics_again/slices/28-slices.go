package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	
	if plan=="planPro" {
		return allMessages[:], nil
	}
	if plan=="planFree" {
		return allMessages[0:2], nil
	}

	return nil, errors.New("Unsupported plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"please, please click here",
		"we beg you to sign up",
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("===================================================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err!=nil {
		fmt.Println("Error: ", err)
		return
	}
	for i:=0; i<len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i==doneAt {
			fmt.Println("user responded")
			break
		}
		if i==len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

func main() {
	test("Goat Matata", 0, planFree)
	test("Dada Ng'ombe", 1, planPro)
	test("Kaka Dubu", 2, planPro)
	test("Mzee Kobe", 3, "no plan")
}