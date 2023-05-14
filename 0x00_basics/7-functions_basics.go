package main

//basics of functions in Go

import (
	"fmt"
	"math"
)

//Use the `func` keyword to declare a function
// func function_name(func_param param_type) { body }
func sayGreeting(greet string) {
	fmt.Printf("Hujambo, %v\n", greet)
}

func sayBye(bye string) {
	fmt.Printf("Kwaheri, %v\n", bye)
}

//functions in Go are first-class citizens, that is, one can store functions in variables,
//pass them to other functions as arguments, and return them from other functions as values

//passing a function as an arg/param...
func cycleNames(n []string, f func(string)) {
	//function `cycleNames` takes in 2 args: `n`, a slice of strings and `f`, a function that
	//accepts one arg of type `string`
	for _, value := range n {
		f(value)
	}

}

//returning a value from a function as a value...
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	animals := []string{"Goat_Matata", "Dada_Ng'ombe", "Kaka_Mbweha", "Mzee_Kobe", "Kaka_Mbwamwitu", "Kaka_Dubu", "Paka_the_Cat", "Kaka_Sungura"}

	fmt.Println("#######--------------------#######")
	fmt.Println("Using a `for` loop to call the greeting function...")
	for _, value := range animals {
		sayGreeting(value)
	}
	fmt.Println("#######--------------------#######")
	fmt.Println("Using a `for` loop to call the goodbye function...")
	for _, value := range animals {
		sayBye(value)
	}
	fmt.Println("#######--------------------#######")
	fmt.Println("Using a function to call the greeting function...")
	cycleNames(animals, sayGreeting)
	fmt.Println("#######--------------------#######")
	fmt.Println("Using a function to call the goodbye function ...")
	cycleNames(animals, sayBye)

	a1 := circleArea(7)
	a2 := circleArea(9.9)

	fmt.Printf("The value of a1 is %0.4f (4dp)\n", a1)
	fmt.Printf("The value of a2 is %0.4f (4dp)\n", a2)

}
