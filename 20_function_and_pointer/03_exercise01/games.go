package main

import (
	"encoding/json"
	"fmt"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

// create a new game with id, name, price, genre
func newGame(id, price int, name, genre string) game {
	return game{
		item{id, name, price},
		genre,
	}
}

// add game to games slice
func addGame(games []game, g game) []game {
	games = append(games, g)
	return games
}

// load the games
func load() (games []game) {
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem: ", err)
		return
	}

	// load the data to games
	// var games []game
	for _, dg := range decoded {
		games = addGame(games, newGame(dg.Id, dg.Price, dg.Name, dg.Genre))
		// games = append(games, game{item{dg.Id, dg.Price, dg.Name}, dg.Genre})

	}

	// games = addGame(games, newGame(1, 50, "god of war", "action adventrue"))
	// games = addGame(games, newGame(2, 40, "x-com 2", "stragegy"))
	// games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))

	return
}

// index game by id
func indexByID(games []game) map[int]game {
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	return byID
}

// function to print game
func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
