package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

// type collection [4]string
type collection []string

func main() {
	// Slice Header:
	// Pointer, Length, Capacity
	// Pointer -> memory location of backing array
	// Length -> number of elements
	// Capacity -> space of backing array, use "cap" to find it.

	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period"}
	change(data)
	s.Show("main's data", data)
	fmt.Printf("main's data slice address: %p\n", &data)

	array := [...]string{"slices", "are", "awesome", "period"}

	// array and slice with same value do not have the same bytes size
	// because: a slice saved only 3 value slice header
	// so slice memory size will be fixed 24 (3*8) bytes
	// however, array size depend on its element numbers
	fmt.Printf("array's size: %d bytes\n", unsafe.Sizeof(array))
	fmt.Printf("slice's size: %d bytes\n", unsafe.Sizeof(data))

}

func change(data collection) {
	// var data collection
	// local data = passed data value
	data[2] = "brilliant!"
	s.Show("change's data", data)
	fmt.Printf("change's data slice address: %p\n", &data)

	// Arguments passed to a function are copied into local variables of the function
	// data won't be same as main data

	// However, when slice is passed into function,
	// it copied only the slice header but not the backing array!!
	// so any change in the argument will also change in global variable
	// Causouly !!!
	// They are difference slice value but with the same backing array!!!

}
