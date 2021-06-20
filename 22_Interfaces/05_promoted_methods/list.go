package main

import "fmt"

// an abstract type, a protocal, a contract.
// no inplementation.
type item interface {
	print()
	discount(ratio float64)
}

type list []item

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery.")
		return
	}
	for _, it := range l {
		// fmt.Printf("(%-10T) ----> ", it)
		it.print()
	}
}

func (l list) discount(ratio float64) {

	for _, it := range l {

		it.discount(ratio)

	}
}
