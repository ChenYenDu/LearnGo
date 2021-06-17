// ---------------------------------------------------------
// EXERCISE: Fix the crash
//
// EXPECTED OUTPUT
//
//   brand: apple
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	// var c *computer  // c was a nil
	c := &computer{} // init with a value
	change(c, "apple")
	fmt.Printf("brand: %s\n", *c.brand) // print the pointed value
}

func change(c *computer, brand string) {
	c.brand = &brand // set the brands address
}
