package main

import "fmt"

func main() {
	// Array length is fixed
	// Slice could grow or shrink during run time

	var numArr [5]int
	fmt.Printf("%#v\n", numArr) // [0, 0, 0, 0, 0]

	var numSlice []int

	fmt.Printf("%#v\n", numSlice) // nil

	// numSlice[0] error!
	// numSlice[1] error!
	fmt.Println(len(numSlice))

}
