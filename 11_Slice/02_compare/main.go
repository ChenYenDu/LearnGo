package main

import "fmt"

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	// var games []string
	// games[0] = "kokemon" error
	// games[1] = "sims" error

	newBooks := [5]string{
		"ulysses", "fire",
	}

	books = newBooks

	games := []string{
		"kokemon",
		"sims",
	}

	newGames := []string{
		"pacman", "doom", "pong",
	}

	// slice type do not contain length so it's ok to assign difference length of slice
	newGames = games

	var ok string

	games = nil
	// ok is still true, why?
	// The following loop is an nil loop
	// Go won't execute a range loop if the given slice is nil

	games = []string{}
	// ok is still true, why?
	// because games is now a initialed slice
	// only uninitiled slice is nil, if initialed will be []

	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	if len(games) != len(newGames) {
		ok = "not "
	}

	// Never check whether a slice is nil or not !!!
	// Instead: Check the length of a slice.

	fmt.Printf("game and newGames are %sequal\n\n", ok)

	/*
		Error!!
		if games == newGames {

		}
		-> a slice can only be compared to a nil value
	*/

	fmt.Printf("books        : %#v\n", books)
	fmt.Printf("games        : %#v\n", games)    // nil
	fmt.Printf("new games    : %#v\n", newGames) // nil

	fmt.Printf("games        : %T\n", games)        // []
	fmt.Printf("games's len  : %d\n", len(games))   // 0
	fmt.Printf("nil?         : %t\n", games == nil) // true

}
