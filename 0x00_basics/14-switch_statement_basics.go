package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
*using the `bill` object in 14-switch_statement_bill.go
*get data from stdin instead of hard-coding it
*use switch statement instead of multiple if-else
*parse floats
* save data to a text file
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
	opt, _ := getInput(" Choose an option:\na - add item\ns - save bill\nt - add tip\n", reader)

	switch opt {
	case "a":
		name, _ := getInput("What's the name of the item?", reader)
		price, _ := getInput("What is the price of the item? (KES)", reader)
		price_float, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}

		b.addItem(name, price_float)
		fmt.Println("item added:", name, price)
		promptOptions(b)

	case "s":
		fmt.Println("Saving the bill...")
		b.save()
		fmt.Println("#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*")
		fmt.Println(b.name, "has been saved")
		fmt.Println("#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*")

	case "t":
		tip, _ := getInput("What is the tip amount? (KES)", reader)

		tip_float, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}

		b.updateTip(tip_float)

		fmt.Println("tip added:", tip)
		promptOptions(b)

	default:
		fmt.Println("Invalid option. Choose 'a', 's' or 't'...")
		promptOptions(b)
	}
}

func main() {
	bill1 := createBill()
	promptOptions(bill1)

	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println(bill1)
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
}
