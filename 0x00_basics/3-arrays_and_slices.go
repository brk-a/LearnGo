package main

import "fmt"

//arrays and slices in Go

func main() {
	//arrays: length is immutable, if that makes any sense
	//method 1: var var_name [arr_len]elem_type = [arr_len]elem_type{elems}
	var numbers [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //arr called `numbers` that can hold 10 elems of type int
	fmt.Println("Arr of ints:", numbers, "Len:", len(numbers))
	numbers[4] = 104
	fmt.Println("Elems are mutable:", numbers, "Len:", len(numbers))

	//method 2: var var_name = [arr_len]elem_type{elems}
	var numbers2 = [10]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //these are of type float, I promise. Check their type.
	fmt.Println("Arr of floats:", numbers2, "Len:", len(numbers2))
	numbers2[7] = 227
	fmt.Println("Elems are mutable:", numbers2, "Len:", len(numbers2))

	//method 3: var_name := [arr_len]{elems}
	strs := [6]string{"Goat_Matata", "Dada_Ng'ombe", "Kaka_Mbweha", "Mzee_Kobe", "Kaka_Sungura"}
	fmt.Println("Arr of strings:", strs, "Len:", len(strs)) //5 elems but len = 6 because strs[5] = null, for now
	strs[2] = "Kaka_Mbwamwitu"
	strs[5] = "Paka_the_Cat"
	fmt.Println("Elems are mutable:", strs, "Len:", len(strs))

	//slices: think python lists
	//len is mutable, if that makes any sense
	//slices are special arrays (deep down, under the bonnet, they are arrays)
	// any of the three methods of c=declaration can be used

	var slice1 = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100} //notice the absence of arr_len
	fmt.Println("Slice of ints:", slice1, "Len:", len(slice1))
	slice1 = append(slice1, 121, 144) //append() returns a new slice
	fmt.Println("Len is mutable:", slice1, "Len:", len(slice1))
	slice1[12] = 1440
	fmt.Println("Elems are mutable too:", slice1, "Len:", len(slice1))

	//BONUS: slice ranges
	// range: part or whole of an array or slice stored in a var of type slice
	//think of slices in R and/or Py; no negative indices. :(
	range1 := strs[0:3] //first index inclusive, second index exclusive
	fmt.Println("Arr of strings:", strs, "Len:", len(strs))
	fmt.Println("strs[0:3]:", range1, len(range1))
	range2 := strs[2:]
	range3 := strs[:3]
	range4 := strs[:]
	fmt.Println("strs[2:]", range2, "Len:", len(range2))
	fmt.Println("strs[:3]", range3, "Len:", len(range3))
	fmt.Println("strs[:]", range4, "Len:", len(range4))
}
