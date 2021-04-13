package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// dynamic-table
	max, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Give me a number.")
		return
	}

	fmt.Printf("%5s", "X")

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
