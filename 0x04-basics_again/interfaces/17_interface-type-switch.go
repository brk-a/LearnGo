package main

import "fmt"

func getExpenseTypeReport(e expenseType) (string, float64) {
	switch v := e.(type) {
	case emailType:
		return v.toAddress, v.cost()
	case smsType:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func (e emailType) cost() float64  {
	if !e.isSubscribed{
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s smsType) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalidType) cost() float64 {
	return 0.0
}

type expenseType interface {
	cost() float64
}

type emailType struct {
isSubscribed bool
body string
toAddress string
}

type smsType struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalidType struct {}

func estimateAnnualCost(e expenseType, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func testType(e expenseType) {
	address, cost := getExpenseTypeReport(e)

	switch e.(type) {
	case emailType:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case smsType:
		fmt.Printf("Report: The smsType going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expenseType")
		fmt.Println("====================================")
	}
}

func main() {
	testType(emailType{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "goatmatata@goatpodcast.com",
	})
	testType(emailType{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "dadangombe@chizicheese.com",
	})
	testType(emailType{
		isSubscribed: false,
		body:         "Wanna catch up later?",
		toAddress:    "mzeekobe@snailmail.com",
	})
	testType(smsType{
		isSubscribed:  false,
		body:          "I'm a Nigerian prince. Please send me your bank info so I can deposit 1000 dollars",
		toPhoneNumber: "+254701234567",
	})
	testType(smsType{
		isSubscribed:  false,
		body:          "I don't need this",
		toPhoneNumber: "+254712345678",
	})
	testType(invalid{})
}