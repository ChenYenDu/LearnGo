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
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
		{title: "minecraft", price: 20},
		{title: "tetris", price: 25},
		{title: "rubik's cube", price: 5},
		{title: "yoda", price: 30},
	}

	store.discount(.5)
	store.print()

}
