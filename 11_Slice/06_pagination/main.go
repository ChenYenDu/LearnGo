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

	const pageSize = 4

	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + 4

		if to > l {
			to = l
		}

		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		s.Show(head, currentPage)
	}

}
