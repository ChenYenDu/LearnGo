package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, byID map[int]game) bool {
	fmt.Println()

	cmd := strings.Fields(input)

	switch cmd[0] {
	case "quit":

		return cmdQuit()

	case "list":
		return cmdList(games)

	case "id":

		return cmdByID(cmd, games, byID)

	case "save":
		return cmdSave(games)

	}
	return true
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdByID(cmd []string, games []game, byID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return true
	}

	printGame(g)
	return true
}

func cmdSave(games []game) bool {

	// load the data to an encodable game values
	var encodeable []jsonGame
	for _, g := range games {
		encodeable = append(encodeable, jsonGame{g.id, g.name, g.genre, g.price})
	}

	out, err := json.MarshalIndent(encodeable, "", "\t")
	if err != nil {
		fmt.Println("Sorry:", err)
		return false
	}
	fmt.Println(string(out))
	return false
}
