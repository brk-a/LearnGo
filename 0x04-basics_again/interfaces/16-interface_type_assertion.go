package main

import "fmt"

func getExpenseReport(e expenses) (string, float64)  {
	em, ok := e.(emailS)
	if ok {
		return em.toAddress, em.cost()
	}

	s, ok := e.(sms)
	if ok {
		return s.phoneNumber, s.cost()
	}

	// invalid, ok := e.(invalid)
	// if ok{
	// 	return "", invalid.cost()
	// }

	return "", 0.0
}

func (e emailS) cost() float64{
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}

	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}

	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expenses interface {
	cost() float64
}

type emailS struct {
	isSubscribed bool
	body string
	toAddress string
}

type sms struct {
	isSubscribed bool
	body string
	phoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expenses, avgMessagesPerYear int) float64 {
	return e.cost() * float64(avgMessagesPerYear)
}

func testAssertion(e expenses) {
	address, cost := getExpenseReport(e)
	
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email to %s will cost %.2f\n", address, cost)
		fmt.Println("===============================================================")
	case sms:
		fmt.Printf("Report: The sms to %s will cost %.2f\n", address, cost)
		fmt.Println("===============================================================")
	default:
		fmt.Printf("Report: Invalid expense")
		fmt.Println("===============================================================")
	}
}

func main() {
	testAssertion(emailS{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "goatmatata@automatic.com",
	})
	testAssertion(emailS{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "kakasungura@automatic.com",
	})
	testAssertion(emailS{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "dadangombe@automatic.com",
	})
	testAssertion(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		phoneNumber: "+254701234567",
	})
	testAssertion(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		phoneNumber: "+254712345678",
	})
	testAssertion(invalid{})
}