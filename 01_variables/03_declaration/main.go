package main

import "fmt"

func main() {
	var speed int
	var heat float64
	var off bool
	var brand string

	/*
		any static type has default values (zero values)
		boolean  -> false
		numerics -> 0
		strings  -> ""
		pointers -> nil
	*/

	fmt.Println(speed)        // default 0
	fmt.Println(heat)         // default 0
	fmt.Println(off)          // default false
	fmt.Printf("%q\n", brand) // default ""

	// multiple declarations
	var (
		speed2 int
		heat2  float64

		off2   bool
		brand2 string
	)

	fmt.Println(speed2)
	fmt.Println(heat2)
	fmt.Println(off2)
	fmt.Println(brand2)

	var speed3, velocity int
	fmt.Println(speed3, velocity)

}
