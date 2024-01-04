package main

import "fmt"

func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}

type emailBill struct {
	costInShillings int
}

func test(bills []emailBill) {
	defer fmt.Println("=====================================================")

	countAdder, costAdder := adder(), adder()
	for _, bill:=range bills {
		fmt.Printf("you have sent %d emails. this costs Kes %d\n", countAdder(1), costAdder(bill.costInShillings))
	} 
}

func main() {
	test([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	test([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	test([]emailBill{
		{743},
		{13},
		{8},
	})
}