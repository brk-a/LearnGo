package main

import "fmt"

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}

	return float64(len(e.body)) * .01
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body string
}

func print(p printer) {
	p.print()
}

func testMult(e expense, p printer) {
	fmt.Print("printing with cost: Kes%.2f ...\n", e.cost())
	p.print()
	fmt.Println("===================================================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body: "Testing...",
	}
	testMult(e, e)

	e = email{
		isSubscribed: false,
		body: "I'd like my money back",
	}
	testMult(e, e)

	e = email{
		isSubscribed: true,
		body: "Can we have a discussion?",
	}
	testMult(e, e)

	e = email{
		isSubscribed: false,
		body: "This meeting could have been a message on slack",
	}
}