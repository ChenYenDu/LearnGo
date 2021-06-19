package main

func main() {
	// var (
	// 	mobydict  = book{title: "moby dick", price: 10}
	// 	minecraft = game{title: "minecraft", price: 20}
	// 	tetris    = game{title: "tetris", price: 25}
	// )

	// // game.print(minecraft)
	// // game.print(tetris)
	// // book.print(mobydict)

	// // GO will automatically take the address when the method has a pointer receiver
	// minecraft.discount(.5) // this equals to (&minecraft).discount(.5)

	// mobydict.print()
	// minecraft.print()
	// tetris.print()

	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}
