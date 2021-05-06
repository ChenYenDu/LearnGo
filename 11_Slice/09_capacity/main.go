package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	// A nil slice don't hava a backing array so its capacity is zero
	s.PrintBacking = true

	var games []string

	// nil slice doens't have a pointer, the capacity field will be zero
	s.Show("games", games)

	games = []string{}
	s.Show("games", games)
	s.Show("another empty", []int{})

	games = []string{"pacman", "mario", "tetris", "doom"}
	s.Show("games", games)

	part := games
	s.Show("part", part)

	part = games[:0]
	s.Show("part", part)
	s.Show("part[:cap]", part[:cap(part)])

	// how to shrink the slice
	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	games = games[len(games):] // this also set slice to an empty slice

	s.Show("games", games)
	s.Show("part", part)
}
