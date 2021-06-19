package main

func main() {
	/*
		Go does not have class:
			Methods can be attached to almost any type
			for example, int, string, float64... etc

		Don't use pointer receivers with these types:
			slice, map, chan, func
	*/

	var (
		// mobydict  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 25}
	)

	tetris.discount(.5)

	var items []*game
	items = append(items, &minecraft, &tetris)

	my := list(items)

	// my = nil

	my.print()
}
