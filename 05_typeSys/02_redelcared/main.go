package main

import (
	"fmt"
	"math"
)

/*
bool, int, int32(rune), uint8(byte), float64
are the most used predeclared type.
And each of them take different size of memory.
*/

func main() {

	// example 1
	var (
		width  uint8 = 255
		height       = 255
	)

	width++ // 256?

	// if width < height {}  -> type error: width is uint8 and height is int
	if int(width) < height {
		fmt.Println("height is greater")
		// this line is printed, why?
		/*
			width is not uint(256) reset to its min value: 0
		*/
	}
	fmt.Printf("width: %d height: %d\n", width, height)

	// example 2
	var n int8 = math.MaxInt8
	fmt.Println("max int8      :", n)   // 127
	fmt.Println("max int8 + 1  :", n+1) // -128, wrap around

	n = math.MinInt8
	fmt.Println("min int8      :", n)   // -128
	fmt.Println("min int8 - 1  :", n-1) // 127, wrap around

	// example 3
	var un uint8
	fmt.Println("max uint8       :", un)   // 0
	fmt.Println("max uint8 - 1   :", un-1) // 255, wrap around

	// example 4
	f := float32(math.MaxFloat32)
	fmt.Println("max float       :", f*1.2) // +Inf
}
