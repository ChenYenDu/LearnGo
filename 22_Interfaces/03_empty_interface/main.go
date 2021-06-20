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

	store := list{
		book{title: "moby dick", price: 10, published: 118281600},
		book{title: "odyssey", price: 15, published: "733622400"},
		book{title: "hobbit", price: 25},
		&game{title: "minecraft", price: 20},
		&game{title: "tetris", price: 25},
		puzzle{title: "rubik's cube", price: 5},
		&toy{title: "yoda", price: 30},
	}

	store.discount(.5)
	store.print()

}
