package main

import "fmt"

func main() {
	i := 144

	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}
}
