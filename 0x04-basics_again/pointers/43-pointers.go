package main

import "fmt"

type Message struct {
	Recipient string
	Text string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

func test(recipient string, text string)  {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=============================================")
}

func main()  {
	test("Goat Matata", "I'll join you on The GOAT Podcast at 0400h every day")
	test("Dada Ng'ombe", "Chizi About Cheese is making an impact out here!")
	test("Mzee Kobe", "How did you become so good at golf?")
}