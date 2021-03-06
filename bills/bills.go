package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//Create new bill
func newBill(nama string) bill {
	b := bill{
		name:  nama,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+": ", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	//total items
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
