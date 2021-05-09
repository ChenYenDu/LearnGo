package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// limit capacity at the same time while you manipulate a slice
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	// odds := nums[:2] // this use the same backing array, also changes nums
	odds := nums[:2:2] // this do limit the capacity to 2, when 2 element were appended, it created a new backing array
	// Use full slice expressions to prevent other code to append more elements to a slice's backing array
	odds = append(odds, 5, 7)

	evens := append(nums[2:4], 6, 8) // the slice's cap is 2 so where it append new elements it create a new backing array

	s.Show("nums", nums)
	s.Show("odds", odds)
	s.Show("evens", evens)
}
