package main

import "fmt"

/**
*`bill` struct and receiver functions
* among others
 */

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... KES %v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ... KES %0.2f\n", "Tip:", b.tip)       // tip
	fs += fmt.Sprintf("%-25v ... KES %0.2f", "Total:", total+b.tip) //total

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip //Go dereferences automatically for structs. Alt: (*b).tip = tip
}

//add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
