package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

main:
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)

		if err != nil {
			// skip non-numerics
			continue
		}

		switch {
		// prime
		case n == 2 || n == 3:
			fmt.Print(n, " ")
			continue

		// not a prime
		case n <= 1 || n%2 == 0 || n%3 == 0:
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			// not a prime
			if n%1 == 0 {
				continue main
			}

			i += w
			w = 6 - w
		}

		fmt.Print(n, " ")
	}

	fmt.Println()
}
