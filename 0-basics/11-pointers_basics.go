package main

import "fmt"

//Pointers, how to create them and how to dereference them

//Default non-pointer types: Str, int, float, bool array, struct
//Default pointer types: slice, map, func
// one can, just like in C, create a pointer for a non-pointer type value
//As ususal, pointers are their own data type and are stored in mem as vars/vals

func updateName(n string) {
	n = "Dada Ng'ombe"
}

func updateNameWithPointer(m *string) {
	*m = "Kaka Mbweha"
}

func main() {
	name := "Goat Matata"
	updateName(name)
	fmt.Println("Print out the value of `name`...")
	fmt.Println(name) //name = "Goat Matata" even after the function `updateName` has run
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	p := &name //pointer to name. stores the addr of the var `name`
	fmt.Println("the address of var `name` is:", &name)
	fmt.Println("the address of the pointer `p` is:", &p)
	fmt.Println("the value of var `name` is:", name)
	fmt.Println("the value of the pointer `p` is:", p)
	fmt.Println("the pointer `p` dereferences to:", *p)
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println("Using pointers to update the values...")
	fmt.Println("value of `name` before calling function w/pointer...", name)
	fmt.Println("Mem addr of `name` befor calling func...", &name)
	updateNameWithPointer(p)
	fmt.Println("value of `name` after calling function w/pointer...", name)
	fmt.Println("Mem addr of `name` after calling func...", &name) //addr is the same.
	//name was updated in-place
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
}

/*
|--name-------|-----p----|
|   0x001     |   0x002  |
|-------------|----------|
|"Goat Matata"|  p0x001  |
|-------------|----------|
*/

/* call the function updateNameWithPointer
|--name-------|-----p----|--m------|
|   0x001     |   0x002  |  0x003  |
|-------------|----------|---------|
|"Kaka Mbweha"|  p0x001  |  p0x001 |
|-------------|----------|---------|
*/
