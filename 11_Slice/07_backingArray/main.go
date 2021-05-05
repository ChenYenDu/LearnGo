package main

import (
	"sort"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// A slice doesn't directly store any elements
	// Slice literal creates a hidden array and return a slice that refers to that array
	// slice -> backing array (real storage which stored separately from the slice)

	// slice is a window of backing array
	// Slices with same shared backing array would change together while alter happened
	// grades := [...]float64{40, 10, 20, 50, 60, 70}
	grades := []float64{40, 10, 20, 50, 60, 70}

	// var newGrades []float64
	// newGrades = append(newGrades, grades...)
	newGrades := append([]float64(nil), grades...)
	// append expect a slice value, but nil is a typeless value
	// so we need to convert nil to a slice value first

	front := newGrades[:3] // break the connection between grades and front
	// front := grades[:3]
	front2 := front[:3]
	front3 := front

	sort.Float64s(front)

	s.PrintBacking = true
	s.MaxPerLine = 7
	s.Show("grades", grades[:])
	s.Show("newGrades", newGrades)
	s.Show("front", front)
	s.Show("front2", front2)
	s.Show("front3", front3)
}
