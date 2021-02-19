package main

import "fmt"

func main() {
	// Printf: prints formatted output
	var brand string = "Google"
	fmt.Printf("%q\n", brand) // return "Google"
	fmt.Printf("%s\n", brand) // return Google

	// Println vs Printf
	var ops = 2350
	var ok = 543
	var fail = 433

	fmt.Println(
		"total:", ops, "success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d success: %d / %d\n", ops, ok, fail,
	)

	// Escape sequences,
	fmt.Println("hi\n\"hi\"")
	fmt.Println("\\")

	// Printing types
	fmt.Printf("%T\n", ops)
	fmt.Printf("%T\n", ok)
	fmt.Printf("%T\n", fail)

	// final presentations
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	fmt.Printf(
		"%v is %v away. Thinks!\n %[2]v kms! %[1]v OMG!\n",
		planet, distance,
	)

	// reformat
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)

	fmt.Printf(
		"%v is %v away. Thinks!\n %[2]v kms! %[1]v OMG!\n",
		planet, distance,
	)

	fmt.Printf(
		"Orbital Period: %f days\n", orbital,
	)

	fmt.Printf(
		"Orbital Period(0 digit): %.0f days\n", orbital,
	)

	fmt.Printf(
		"Orbital Period(2 digit): %.2f days\n", orbital,
	)

}
