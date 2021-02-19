package main

import "fmt"

// when you need a package scope variable, use normal declatation
var version string

func main() {
	// If you don't know the initial value use normal declaration
	var score int

	// When you want to group variables together for greater readability, use normal declaration
	var (
		// related
		video string

		// closely related
		duration int
		current  int
	)

	fmt.Println(version, score, video, duration, current)

	// We mostly prefer using "Short declaration"
	// If you knoe the initital value, use short declaration
	width, height := 100, 50

	// To keep the code concise, use short declaration
	// For redeclaration, use short declaration
	width, color := 200, "red"

	fmt.Println(width, height, color)

}
