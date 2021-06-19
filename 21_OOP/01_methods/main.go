package main

func main() {
	mobydict := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 25,
	}

	mobydict.print()
	minecraft.print()
	tetris.print()
}
