package main

import (
	"fmt"
	"strings"
)

func main() {
	// args := os.Args[1:]

	// for i := 0; i < len(args); i++ {
	// 	fmt.Printf("%q\n", args[i])
	// }

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}
}
