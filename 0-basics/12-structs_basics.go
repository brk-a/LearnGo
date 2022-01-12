package main

import "fmt"

/**
*struct and custom data types.
*Structs are similar to their cousins in C
*think python classes, more or less
 */

//struct is a blueprint/mold/cookie cutter that represents the data
// that is desired.
//a struct may contain any number of different [basic] data types eg.
//str, int, arr, slice, map, bool etc

func main() {
	bill1 := newBill("Goat Matata")
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Printf("%v's bill:\n", bill1.name)
	//go run 12* .go on cmd line
	fmt.Println(bill1)
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println("Using the `format` receiver function...")
	fmt.Println(bill1.format())
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	bill2 := newBill("Dada Ng'ombe")
	bill2.updateTip(3.14)
	bill2.addItem("Lucerne, fat free", 20.48)
	bill2.addItem("Napier, classic", 4.25)
	bill2.addItem("Water, zero sugar", 2.13)
	fmt.Printf("%v's bill:\n", bill2.name)
	fmt.Println(bill2)
	fmt.Println(bill2.format())
}
