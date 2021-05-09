package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// `copy` copies elements of a slice to another slice
	// E.g. copy( `destination slice`, `source slice`)
	// copy only use the smallest length of both destination slice and source slice
	s.PrintBacking = true
	s.MaxPerLine = 10

	data := []float64{10, 25, 30, 50}

	// Method 1:You can do it with loop
	// newData := []float64{99, 100}
	// for i := range newData {
	// 	data[i] = newData[i]
	// }

	// Method 2: copy
	copy(data, []float64{99, 100})
	// copy(data, []float64{10, 5, 15, 0, 20})

	// saved := make([]float64, len(data))
	// copy(saved, data)
	// Same as:
	saved := append([]float64(nil), data...)
	data[0] = 0

	s.Show("Probabilities (saved)", saved)
	s.Show("Probabilities (data)", data)

	fmt.Printf(
		"Is it gonna rain? %.f%% chance.\n",
		(data[0]+data[1]+data[2]+data[3])/float64(len(data)),
	)

}
