package main

import "fmt"

func main() {
	blue := [3]int{6, 9, 3}

	red := blue
	/*
		Error! Different types of arrays are not assignable
		red := [2]int{3, 5}
		red = blue
	*/

	blue[0] = 10 // red[0] is still 6

	fmt.Printf("blue books:  %v\n", blue)
	fmt.Printf("red books:   %v\n", red)

	prev := [3]string{
		"Kafka's Revenge", "Stay Golden", "Everythingship",
	}

	// books := prev
	var books [4]string

	for i, b := range prev {
		books[i] += b + " 2nd Ed."
	}
	books[3] = "Awesome"

	fmt.Printf("last year: \n%#v\n", prev)
	fmt.Printf("this year: \n%#v\n", books)

}
