package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Args variable belongs to the os package
	*/
	// fmt.Printf("%#v\n", os.Args)

	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("1st argument:", os.Args[1])
	// fmt.Println("2nd argument:", os.Args[2])
	// fmt.Println("3rd argument:", os.Args[3])

	// fmt.Println("Number of items inside os.Args:", len(os.Args))

	// var name string
	// name = os.Args[1]
	name := os.Args[1]
	fmt.Println("Hello great", name, "!")

	name, age := "gdanalf", 2019
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}
