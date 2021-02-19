package main

import "fmt"

func main() {

	speed := 100 // speed is int
	force := 2.5 // force is float64

	// error: mismatch type operation
	// speed = speed * force (error !!!!)

	// force is converted to 2 from 2.5
	speed = speed * int(force)
	fmt.Println(speed)
	fmt.Println(force, int(force))

	// speed is converted to 200.0
	speed = int(float64(speed) * force)
	fmt.Println(speed)

	// int & int32 are different types
	var apple int
	var orange int32

	// apple = orange (compiler error)
	apple = int(orange)
	fmt.Println(apple)

	// isDelicious := bool(orange)

	orange = 65
	color := string(orange) // return "A"
	fmt.Println(color)
	// fmt.Println(string(65.0)) error!
	fmt.Println(string([]byte{104, 105})) // return hi, 104 -> h, 105 -> i

}
