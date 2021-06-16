package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(".... Arrays")
	arrays()

	fmt.Println(".... Slices")
	slices()
	// slice already contains a pointer to its backing array,
	// you don't need to use a pointer to modify element in slice
	// passing slice to function are not comman,
	// so `DO NOT USE POINTER WITH SLICES`

	fmt.Println(".... Maps")
	maps()
	// map value is already a pointer
	// do not use pointer want working with map in function

	fmt.Println(".... Structs")
	structs()
	// go automatically use pointer value want passing pointer of struct to a function

}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)

	// fmt.Printf("%+v\n", myHouse)
	fmt.Printf("struct house    : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)

	// fmt.Printf("%+v\n", myHouse)
	fmt.Printf("struct house    : %p %+v\n", &myHouse, myHouse)
}

func addRoomPtr(h *house) {
	h.rooms++ // same as: (*h)++
	fmt.Printf("addRoomPtr house    : %p %+v\n", h, h)

	fmt.Printf("&myHouse.name : %p\n", &h.name)
	fmt.Printf("&myHouse.rooms: %p\n", &h.rooms)

}
func addRoom(h house) {
	h.rooms++
	fmt.Printf("addRoom house    : %p %+v\n", &h, h)

}

// ..................................................

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

// ..................................................

func slices() {
	dirs := []string{"up", "right", "down", "left"}

	up(dirs)
	fmt.Printf("slice dirs			: %p %q\n", &dirs, dirs)
	upPtr(&dirs)
}

func upPtr(list *[]string) {
	// slice do not automatically work in pointer
	// so you need to assign a new variable for it's value
	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "HEISEN BUG")
	// this will end up having bugs, by creating a new slice header with new pointer

	fmt.Printf("upPtr list					: %p %q\n", list, list)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	list = append(list, "HEISEN BUG")
	// this will end up having bugs, by creating a new slice header with new pointer

	fmt.Printf("up list					: %p %q\n", &list, list)
}

// ..................................................
func arrays() {
	nums := [...]int{1, 2, 3}
	fmt.Printf("arrays nums			: %p\n", &nums)

	incr(nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("incr nums			: %p\n", &nums)

	for i := range nums {
		nums[i]++
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums			: %p\n", &nums)

	for i := range nums {
		nums[i]++
	}
}
