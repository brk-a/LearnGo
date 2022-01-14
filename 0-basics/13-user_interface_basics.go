package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
*using the `bill` object in 13-user_interface_bill.go
*get data from stdin instead of hard-coding it
 */

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new name for a bill: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill: ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(" Choose an option:\na - add item\ns - save item\nt - add tip", reader)

	if opt == "a" {
		//code here
	} else if opt == "s" {
		//code here
	} else if opt == "t" {
		//code here
	} else {
		//code  here
	} // use `switch` statement in place of multiple `if-else`
}

func main() {
	bill1 := createBill()
	promptOptions(bill1)

	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println(bill1)
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
}
