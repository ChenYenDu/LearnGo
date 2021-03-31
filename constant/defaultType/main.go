package main

import "fmt"

func main() {
	const min int32 = 1
	max := 5 + min

	fmt.Printf("Type of max is %T\n", max) // int32

	const unMin = 1
	unMax := 5 + unMin

	fmt.Printf("Type of unMax is %T\n", unMax) // int

	/*
		Default type
	*/
	i := 42      // int
	f := 3.14    // float64
	b := true    // bool
	s := "Hello" // string
	r := 'A'     // rune

	fmt.Printf(
		"%T %T %T %T %T\n", i, f, b, s, r,
	)

	j := 42 + 3.14 // float64
	fmt.Printf(
		"Type of j %T\n", j,
	)

	// t := true + "Hello" (ERROR!)
}
