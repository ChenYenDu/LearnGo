package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// `make` initializes and returns a slice with the given length and capacity
	// E.g. make([]int, 3, 5) -> length: 3 and capacity: 5

	s.PrintBacking = true
	s.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}

	// Prevent overwriting the "tasks" by using "upTasks"
	// var upTasks []string
	// for _, task := range tasks {
	// 	upTasks = append(upTasks, strings.ToUpper(task))
	// 	s.Show("upTasks", upTasks)
	// }
	// This pratice creates 2 element due to the capacity lackness
	upTasks := make([]string, 0, len(tasks)) // 0 length, 3 capacity
	s.Show("upTasks", upTasks)

	// This time only use a backing array
	for _, task := range tasks {
		// upTasks[i] = strings.ToUpper(task)
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}

}
