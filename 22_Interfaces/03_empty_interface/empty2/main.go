package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	var any interface{}
	any = nums

	// _ = len(any); won't work
	_ = len(any.([]int))

	var many []interface{}
	// many = nums
	for _, n := range nums {
		many = append(many, n)
	}
	fmt.Println(many)
}
