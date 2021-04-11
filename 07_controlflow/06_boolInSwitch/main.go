package main

import "fmt"

func main() {

	i := 10

	switch { // switch condition default: true

	case i > 0:
		// positive
		fmt.Println("positive")

	case i < 0:
		// negative
		fmt.Println("negative")
	default:
		// zero
		fmt.Println("zero")

	}
}
