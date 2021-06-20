package main

import "fmt"

// an abstract type, a protocal, a contract.
// no inplementation.
type printer interface {
	print()
}

type list []printer

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
	type discounter interface {
		discount(float64)
	}
	for _, it := range l {
		// instead of checking a single type,
		// g, isGame := it.(*game)
		// use a interface checking work even better
		// g, isGame := it.(discounter)

		// if !isGame {
		// 	continue
		// }

		// g.discount(ratio)
		// For cleaner code
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}
