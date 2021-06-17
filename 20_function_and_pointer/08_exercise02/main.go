// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.

// ---------------------------------------------------------

package main

import "fmt"

func main() {
	//  1- Swap the values through a func
	//
	//     a- Declare two float variables
	//
	//     b- Declare a func that can swap the variables' values
	//        through their memory addresses
	//
	//     c- Pass the variables' addresses to the func
	//
	//     d- Print the variables
	a, b := 3.14, 6.28
	swap(&a, &b)
	fmt.Printf("a : %g - b : %g\n", a, b)

	//  2- Swap the addresses of the float pointers through a func
	//
	//     a- Declare two float pointer variables and,
	//        assign them the addresses of float variables
	//
	//     b- Declare a func that can swap the addresses
	//        of two float pointers
	//
	//     c- Pass the pointer variables to the func
	//
	//     d- Print the addresses, and values that are
	//        pointed by the pointer variables
	pa, pb := &a, &b
	pa, pb = swapAddr(pa, pb)
	fmt.Printf("pa: %p - pb: %p\n", pa, pb)
	fmt.Printf("pa: %g - pb: %g\n", *pa, *pb)

}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func swapAddr(a, b *float64) (*float64, *float64) {
	return b, a
}
