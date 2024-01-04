package main

import "fmt"

func printReports(messages []string) {
	for _, message:= range messages {
		printCostReport(func(i string) int {
			return len(i) * 2
		}, message)
	}
}

func test(messages []string) {
	defer fmt.Println("========================================================")
	printReports(messages)
}

func main() {
	test([]string{
		"Here's Goat Matata",
		"He owns and runs 'The GOAT Podcast'",
	})
	test([]string{
		"Here's Dada Ng'ombe",
		"She owns and runs 'Chizi About Cheese'",
	})
	test([]string{
		"Here's Paka the Cat",
		"He owns and runs 'Catnip & Chill'",
	})
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: [%s] Cost: Kes %v`, message, cost)
	fmt.Println()
}