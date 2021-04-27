package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// append(slice, newElement)
	nums := []int{1, 2, 3}
	fmt.Println("first length     :", len(nums))

	nums = append(nums, 4, 9)
	fmt.Println("second length    :", len(nums))

	tens := []int{12, 13}
	nums = append(nums, tens...)
	fmt.Println("third length     :", len(nums))

	// create a Todo list
	var todo []string

	s.Show("todo", todo) // (len:0 cap:0 ptr:0)

	todo = append(todo, "sing")
	s.Show("todo", todo) // (len:1 cap:1 ptr:XXXX)

	todo = append(todo, "run", "code", "play")
	s.Show("todo", todo) // (len:4 cap:4, ptr:XXXX)
}
