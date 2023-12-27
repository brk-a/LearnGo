package main

import "fmt"

type message struct {
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number string
}

func canSendMessage(mToSend message) bool {
	if mToSend.sender.name=="" {
		return false
	}
	if mToSend.recipient.name=="" {
		return false
	}
	if mToSend.sender.number=="" {
		return false
	}
	if mToSend.recipient.number=="" {
		return false
	}

	return true
}

func testCanSendMessage(mToSend message)  {
	fmt.Printf(
		`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient,
		mToSend.recipient.number,
	)
	fmt.Println()

	if canSendMessage(mToSend) {
		fmt.Println("... sent.")
	} else {
		fmt.Println("... cannot send message.")
	}

	fmt.Println("========================================================")
}

func main() {
	testCanSendMessage(message{
		message: "golf with mzee kobe",
		sender: user{
			name: "Bi Buibui",
			number: "0701234567",
		},
		recipient: user{
			name: "F Njakai",
			number: "0712345678",
		},
	})

	testCanSendMessage(message{
		message: "the GOAT podcast with goat matata",
		sender: user{
			name: "Bi Buibui",
			number: "0701234567",
		},
		recipient: user{
			name: "F Njakai",
			number: "",
		},
	})

	testCanSendMessage(message{
		message: "fly fishing with kaka dubu",
		sender: user{
			number: "0701234567",
		},
		recipient: user{
			name: "F Njakai",
			number: "0712345678",
		},
	})

	testCanSendMessage(message{
		message: "cheese tasting with data ng'ombe",
		sender: user{
			name: "",
			number: "0701234567",
		},
		recipient: user{
			name: "F Njakai",
			number: "0712345678",
		},
	})
}