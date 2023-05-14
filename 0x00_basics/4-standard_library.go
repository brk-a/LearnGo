package main

import (
	"fmt"
	"sort"
	"strings"
)

//Go, from the box, has a lightweight set of features and functions
//Enter the std lib

func main() {
	greeting := "Hello, my people!"

	//strings package
	fmt.Println(strings.Contains(greeting, "people"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) //returns a new string
	fmt.Println("Original string is not changed:", greeting)
	fmt.Println("All upper:", strings.ToUpper(greeting))
	fmt.Println("All lower:", strings.ToLower(greeting))
	fmt.Println("Index of first occurrence of `ll` in the original string:", strings.Index(greeting, "ll"))
	fmt.Println("Split a sting to an array by the space:", strings.Split(greeting, " "))

	//sort package
	squares := []int{0, 81, 4, 49, 16, 25, 64, 36, 9, 1}
	index := sort.SearchInts(squares, 16)
	fmt.Println("Original slice:", squares)
	fmt.Println("Index of `16` is:", index) //16 is not at position 3!! WTF!?
	sort.Ints(squares)                      //returns the same `squares` slice, not a copy
	fmt.Println("Using sort.Ints():", squares)
	fmt.Println("sort.Ints() alters the original slice:", squares)
	index = sort.SearchInts(squares, 16)
	fmt.Println("Index of `16` is:", index)

	animals := []string{"Goat_Matata", "Dada_Ng'ombe", "Kaka_Mbweha", "Mzee_Kobe", "Paka_the_Cat", "Kaka_Mbwamwitu"}
	idx := sort.SearchStrings(animals, "Kaka_Mbweha")
	fmt.Println("Original slice:", animals)
	fmt.Println("Index of `Kaka_Mbweha` is:", idx) //somehow, it got the index right. Interesting...
	sort.Strings(animals)                          //returns the same `squares` slice, not a copy
	fmt.Println("Using sort.Strings():", animals)
	fmt.Println("sort.Strings() alters the original slice:", animals)
	idx = sort.SearchStrings(animals, "Kaka_Mbweha")
	fmt.Println("Index of `Kaka_Mbweha` is:", idx)
}
