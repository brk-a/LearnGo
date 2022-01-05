package main

//For Loops in Go

import "fmt"

func main() {
	// The `while` keyword does not exist in Go (as at 05-Jan-2022)
	//Manipulate the `for` keyword to achieve the same effect
	x := 0
	fmt.Println("`for` acting as `while`...")
	for x < 5 {
		fmt.Println("the value of x is:", x)
		x++
	}

	//same thing, using `for` acting as itself
	fmt.Println("`for` as `for`...")
	for i := 0; i < 5; i++ {
		fmt.Println("the value of i is:", i)
	}

	//traversing a slice...
	animals := []string{"Goat_Matata", "Dada_Ng'ombe", "Kaka_Mbweha", "Mzee_Kobe", "Kaka_Mbwamwitu", "Kaka_Dubu", "Paka_the_Cat"}
	fmt.Println("Slice of animals using `for` as `for`...")
	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	//traversing a slice: using `for` acting as `for - in` ...
	fmt.Println("`for` acting as `for - in `...")
	for idx, val := range animals { // idx: index of elem in slice. val: elem of slice.
		fmt.Printf("the value at index %v is %v\n", idx, val)
	}

	//what if I don't need one of those? Say, I don't need idx?
	//Eeasy...!
	fmt.Println("What if, say, I don't need index? Eeasy! use `_` in place of idx...")
	for _, val := range animals {
		fmt.Printf("The value is %v\n", val)
	}
}
