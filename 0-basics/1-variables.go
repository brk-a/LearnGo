package main

import "fmt"

func main() {
	//method 1: var var_name var_type = value
	var name1 string = "Goat Matata"

	//method 2: var var_name = value
	//Go infers the type based on the value
	var name2 = "Dada Ng'ombe"
	var number1 = 123
	var boolean1 = false
	var number2 = 123.45

	//method 3: var var_name var_type
	//One can then assign a value when needed
	var name3 string  //default: empty string
	var boolean2 bool //default: false

	fmt.Println(name1, name2, number1, boolean1, number2, name3, boolean2)

	//update name3 and boolean2
	name3 = "Kaka Mbweha"
	boolean2 = true

	//method 4: var_name := value
	//Go infers the type based on the value
	//Use this to initialise, not update, a variable
	//Use this inside, not outside, a function
	number3 := 456
	number4 := 456.78
	name4 := "Mzee Kobe"
	boolean3 := false

	fmt.Println(name1, name2, number1, boolean1, number2, name3, boolean2, number3, number4, name4, boolean3)
}
