package main

import (
	"fmt"
	"strings"
)

func main() {

	// for _, v := range os.Args[1:] {
	// 	fmt.Printf("%q\n", v)
	// }

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	for i, v := range words {
		fmt.Printf("#%-2d: %q\n",
			i+1, v,
		)
	}
}
