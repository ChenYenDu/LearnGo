package main

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	// showing you how pointer, length, capacity change
	// append set capacity to doubles the length when the slice is full
	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)

		time.Sleep(time.Second / 4)
	}

	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		c := float64(cap(ages))

		if c != oldCap {
			fmt.Printf(
				"len:%-10d cap:%-10g growth:%.2f\n",
				len(ages), c, c/oldCap,
			)
		}
		oldCap = c
	}

	// append adds less element to slice,
	// when there are to much elements (> 1024),
	// So do not permaturelly optimize,
	// trust the append function.
}
