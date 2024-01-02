package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costCustomer, err := sendSMS(msgToCustomer)
	if err!=nil {
		return 0.0, err
	}

	costSpouse, err := sendSMS(msgToSpouse)
	if err!=nil {
		return 0.0, err
	}

	return costCustomer + costSpouse, nil
}

func sendSMS(message string) (float64, error)  {
	const maxTextLen = 25
	const costPerChar = .0002

	if len(message)>maxTextLen {
		return 0.0, fmt.Errorf("Cannot send texts over %v characters long", maxTextLen)
	}

	return costPerChar * float64(len(message)), nil
}

func testSend(msgToCustomer, msgToSpouse string)  {
	defer fmt.Println("============================================")
	fmt.Println("Message to customer: ", msgToCustomer)
	fmt.Println("Message to spouse: ", msgToSpouse)

	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err!=nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Total cost: Kes %.4f\n", totalCost)
}

func main() {
	testSend(
		"Thanks for coming in to our flower shop today!",
		"i ::heart:: you",
	)
	testSend(
		"Thanks for joining us!",
		"Have a good day.",
	)
	testSend(
		"Thank you.",
		"Enjoy!",
	)
	testSend(
		"We loved having you in!",
		"happy birthday",
	)
}