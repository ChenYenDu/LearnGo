package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// Sliceable -> Slice, Array, String
	// newSlice = Sliceable[Start:Stop]
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galega", "frogger", "astroid", "simicty",
		"metriod", "defender", "raymen", "tempast",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("items", items)

	top3 := items[:3]
	s.Show("top3", top3)

	l := len(items)

	last4 := items[l-4:]
	s.Show("last4", last4)

	// slicing a slice is called: "Reslicing"
	mid := last4[1:3]
	s.Show("mid", mid)

	// slicing return a slice value, and indexing return a single value
	fmt.Printf("slicing:     %T %[1]q\n", items[2:3]) // []string ["tetris"]
	fmt.Printf("indexing:    %T %[1]q\n", items[2])   // string "tetris"
}
