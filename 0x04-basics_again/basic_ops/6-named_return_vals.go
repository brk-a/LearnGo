package main

import (
	"fmt"
)

func main() {
	test(4)
	test(15)
	test(18)
	test(30)
	test(99)
}

func yearsUntilEvents(age int) (
	yearsUntilAdult int,
	yearsUntilDrinking int,
	yearsUntilCarRental int,
	) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult<0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 18 - age
	if yearsUntilDrinking<0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 20 - age
	if yearsUntilCarRental<0 {
		yearsUntilCarRental = 0
	}

	return
}

func test(age int)  {
	fmt.Println("Age: %v", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)

	fmt.Println("You are an adult in %v years", yearsUntilAdult)
	fmt.Println("You can drink in %v years", yearsUntilDrinking)
	fmt.Println("You can rent a car in %v years", yearsUntilCarRental)
	fmt.Println("================================================================")
}