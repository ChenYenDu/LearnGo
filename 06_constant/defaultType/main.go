package main

import (
	"fmt"
	"time"
)

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

	// const (
	// 	Nanosecond  time.Duration = 1
	// 	Microsecond               = Nanosecond * 1000
	// 	Millisecond               = Microsecond * 1000
	// 	Second                    = Millisecond * 1000
	// 	Minute                    = Second * 60
	// 	Hour                      = Minute * 60
	// )

	t := time.Second * 10
	fmt.Println(t)

	it := 10
	// t = time.Second * i (Error)
	t = time.Second * time.Duration(it)
	fmt.Println(t)

	const ct = 10
	t = time.Second * ct
	fmt.Println(t)
}
