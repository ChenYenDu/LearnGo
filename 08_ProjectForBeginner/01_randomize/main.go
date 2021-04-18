package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 5

func main() {

	// set seed with time
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Pick a number")
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You Win!")
			return
		}

	}
	fmt.Println("You lose!")

	// for n := 0; n != guess; {
	// 	n = rand.Intn(guess + 1)
	// 	fmt.Printf("%d ", n)
	// }
	// fmt.Println()

}
