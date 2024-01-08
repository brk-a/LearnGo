package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email:=range emails {
		emailsToSend <- email
	}

	return emailsToSend
}

func sendEmails(batchSize int, ch chan string)  {
	for i:=0; i<batchSize; i++ {
		email := <- ch
		fmt.Println("Sending emails: ", email)
	}
}

func test(emails ...string)  {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("sending emails... ")
	sendEmails(len(emails), ch)
	fmt.Println("==================================================")
}

func main() {
	test("Hi, F Njakai. Tell Bethel I said hi. Goat Matata.", "Wagwaan, man. Goat Matata here.")
	test("I find that hard to believe. F Njakai", "When? I don't know if I can. F Njakai", "What time are you thinking of? F Njakai.")
	test("The calves say hi. Dada Ng'ombe", "Yes, it's tomorrow. So we're good, innit. Dada Ng'ombe.", "Cool. See you then! Dada Ng'ombe.")
}