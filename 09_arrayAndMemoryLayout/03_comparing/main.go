package main

import "fmt"

func main() {
	var (
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}
		// red = [5]int{6, 9, 3} type are different could not compare
	)

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}
