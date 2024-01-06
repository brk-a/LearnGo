package main

import "fmt"

func (e *email) setMessage(newMessage string)  {
	e.message = newMessage
}

type email struct {
	message string
	fromAddress string
	toAddress string
}

func test(e *email, newMessage string)  {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("============================================")
}

func (e email) print() {
	fmt.Println("message: ", e.message)
	fmt.Println("fromAddress: ", e.fromAddress)
	fmt.Println("toAddress: ", e.toAddress)
}

func main()  {
	test(&email{
		message: "first try...",
		fromAddress: "goatmatata@example.co.ke",
		toAddress: "dadangombe@example.co.ke",
	}, "second try...")
}