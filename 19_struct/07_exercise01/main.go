package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {

	type item struct {
		id, price int
		name      string
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

	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem: ", err)
		return
	}

	// load the data to games
	var games []game
	for _, dg := range decoded {
		games = append(games, game{item{dg.Id, dg.Price, dg.Name}, dg.Genre})

	}

	// games := []game{
	// 	{item{id: 1, name: "god of war", price: 50}, "action adventure"},
	// 	{item{id: 2, name: "x-com 2", price: 30}, "strategy"},
	// 	{item{id: 3, name: "minecraft", price: 20}, "sandbox"},
	// }

	byIDs := make(map[int]game)

	for _, game := range games {
		byIDs[game.id] = game
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> list : list all games 
> id N : query a game by id 
> save : export the data to json and quits
> quit : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return
		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		case "id":
			if len(cmd) != 2 {
				fmt.Println("Please enter an id.")
				continue
			}

			id, err := strconv.Atoi(cmd[1])

			if err != nil {
				fmt.Println("Please enter a valid id.", cmd[1])
				continue
			}

			g, ok := byIDs[id]

			if !ok {
				fmt.Printf("Sorry, I can't find game with id: %d", id)
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)

		case "save":

			// load the data to an encodable game values
			var encodeable []jsonGame
			for _, g := range games {
				encodeable = append(encodeable, jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodeable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				return
			}
			fmt.Println(string(out))
			return

		}
	}

}
