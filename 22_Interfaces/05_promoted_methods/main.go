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

	store := list{
		&book{product{"moby dick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 25}},
		&puzzle{product{"rubik's cube", 5}},
		&toy{product{"yoda", 30}},
	}

	store.discount(.5)
	store.print()

	t := &toy{product{"yoda", 150}}
	fmt.Printf("%#v\n", t)

	b := &book{product{"moby dick", 10}, 118281600}
	fmt.Printf("%#v\n", b)
}
