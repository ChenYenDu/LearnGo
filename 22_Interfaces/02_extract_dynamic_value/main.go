package main

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
		yoda      = toy{title: "yoda", price: 30}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydict, rubik, &yoda)
	// store.print()

	// zero interface value is nil
	var p printer

	p = mobydict
	p = rubik
	p = &minecraft

	p = &tetris
	tetris.discount(.5)
	p.print()
	// p.discount(.5)
	// you cannot use this because a interface do not container discount method
	// it belong to dynamic type

	store.discount(.5)
	store.print()

}
