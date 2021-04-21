package main

import "fmt"

func main() {
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
	}
	fmt.Printf("books   : %q\n", books) // automatically add "" to unfilled space

	cars := [...]string{"Tyota", "Benz"} // [...] automatically detect length of array
	fmt.Printf("cars    : %q\n", cars)
}
