package main

import "fmt"

func main() {
	// empty interface says nothing
	// interface {}
	var any interface{}

	// you can store any type of value in an empty interface value
	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "hello"
	any = 3
	// but, you cannot directly use the dynamic value of an empty interface value
	// any = any * 2; won't work
	any = any.(int) * 2
	fmt.Println(any)
}
