package main

import "fmt"

func main() {

	/*
		Interface is a `Protocol` - a Contract
		Only describe `the expected behavior`

		e.g.
		Many Devices -> PowerDrawer Interface (Draw(power int)) -> Socket

		-> i (socket) need a draw method but i don't care where it comes from

		``` Bigger the interface the weaker the abstraction``` -- Rob Pike
	*/

	var (
		mobydict  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 25}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydict, rubik)
	store.print()

	// interface values are compariable
	fmt.Println(store[0] == &minecraft)
	fmt.Println(store[3] == rubik)

}
