package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		// onley: a variable
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		// only: a, n and err variables
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// all the variables in the if statement
		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}

	fmt.Printf("n is %d. 👻 👻 👻 - you've been shadowed ;", n)
}
