package main

import "fmt"

// IOTA is a built-on constant generator which generates ever increasing numbers
func main() {
	// const (
	// 	monday    = 0
	// 	tuesday   = 1
	// 	wednesday = 2
	// 	thursday  = 3
	// 	friday    = 4
	// 	saturday  = 5
	// 	sunday    = 6
	// )

	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	// Blank-Identifier
	// const (
	// 	EST = -5
	// 	MST = -7
	// 	PST = -8
	// )

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(EST, MST, PST)
}
