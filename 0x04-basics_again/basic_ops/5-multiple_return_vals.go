package main

import (
	"fmt"
)

func main()  {
	fname, lname := getNames()
	first, _ := getNames()
	_, last := getNames()

	fmt.Println("Hello, %s %s.", fname, lname)
	fmt.Println("Welcome to Textio, %s.", first)
	fmt.Println("Your last name is %s.", last)
}

func getNames() (string, string) {
	return "F", "Njakai"
}