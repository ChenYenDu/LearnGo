package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// percedence -> the order of how the code work
	celsius := 35.

	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g ˚C is %g KF\n", celsius, fahrenheit)

	// incedec statement
	var n int
	fmt.Println(n)
	n = n + 1
	fmt.Println(n)
	n += 1
	fmt.Println(n)
	n++
	fmt.Println(n)

	n = 10
	fmt.Println(n)
	n = n - 1
	fmt.Println(n)
	n -= 1
	fmt.Println(n)
	n--
	fmt.Println(n)

	// example
	var counter int
	counter++
	counter++
	counter++
	counter--
	fmt.Printf("There ara %d line(s) in the file\n", counter)

	// exercise: count area
	width := 100
	height := 25

	area := width * height
	fmt.Printf("The area is %d feet-square\n", area)

	area -= 10
	area += 10
	area *= 2
	area /= 2
	fmt.Printf("The area is %d now\n", area)
	fmt.Printf("%T", area)

	// exercise: feet to meter
	// strconv: allows you to convert basic values from/to a string value
	arg := os.Args[1] // user input, first element is the path where package work

	feet, _ := strconv.ParseFloat(arg, 64)
	fmt.Printf("Type of feet: %T\n", feet)

	meters := feet * 0.3048
	fmt.Printf("%g feet equals to %g meters\n", feet, meters)

	// exercise: Celsius to Fahrenheit
	arg2 := os.Args[2]

	degree, _ := strconv.ParseFloat(arg2, 64)

	faDegree := degree*1.8 + 32

	fmt.Printf("%g celius equals to %g fahrenheit\n", degree, faDegree)

}
