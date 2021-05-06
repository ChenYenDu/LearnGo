package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	var nums []int
	s.Show("no backing array", nums)

	nums = append(nums, 1, 3)
	s.Show("allocates", nums) // printer has been changed, and length == capacity

	nums = append(nums, 2)
	s.Show("free capacity", nums) // pointer is changed again, length = 3, capacity=4

	nums = append(nums, 4)
	s.Show("no allocation", nums) // pointer is not change, length = 4, capacity = 4

	nums = append(nums, nums[2:]...)
	s.Show("nums <- nums[2:]", nums)

	// append function add element to the end of the slice depend on its `length`
	nums = append(nums[:2], 7, 9)
	s.Show("nums[:2] <- 7,9", nums) // pointer, length, capacity all still same but the last two elements is changed

	// get 2, 4 back to element
	nums = nums[:6]
	s.Show("nums: extend", nums)
}
