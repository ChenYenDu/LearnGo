package main

import "fmt"

func main() {
	// var safe bool = true
	safe := true
	fmt.Println(safe)

	var name string = "Carl"

	var inScientist bool = true

	var age int = 62

	var degree float64 = 5.

	fmt.Println(name, inScientist, age, degree)

	// multiple short declatation
	safe2, speed := true, 50

	fmt.Println(safe2, speed)

	// redeclaration
	safe = false
	fmt.Println("redeclaration", safe)
}
