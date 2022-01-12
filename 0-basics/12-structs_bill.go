package main

import "fmt"

/**
*structs and receiver functions to be used
*in 12-structs_basics.go
 */

//receiver functions: think python class methods

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//receiver function to format bill
// func (var_name, struct_name) func_name(optional_params) return_type {body}
func (b *bill) format() string {
	//the function `format` is tied to the `bill` type and returns a `string` type
	//the `b` before bill is, simply, a var to help one access `bill` via a pointer
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list of items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... KES %v\n", k+":", v)
		total += v
	}

	//tip
	fs += fmt.Sprintf("%-25v ... KES %0.2f\n", "Tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ... KES %0.2f", "Total:", total+b.tip)

	return fs
}

//update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip //Go dereferences automatically for structs. Alt: (*b).tip = tip
}

//add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
