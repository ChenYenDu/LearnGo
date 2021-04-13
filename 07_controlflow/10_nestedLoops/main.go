package main

import "fmt"

func main() {
	const max = 5

	fmt.Printf("%5s", "X")

	for i := 0; i <= 5; i++ {
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
