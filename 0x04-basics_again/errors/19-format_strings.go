package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs Kes %.2f to be sent to '%v' cannot be sent",
		cost,
		recipient,
	)
}

// don't edit below this line

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+254 (701) 234 567")
	test(2.1, "+254 (712) 345 678")
	test(32.1, "+254 (723) 456 789")
	test(14.4, "+254 (734) 567 890")
}