package main

import "fmt"

func main() {
	// constants belong to compile time
	// constants may or may not have type
	const meters = 100

	cm := 100
	m := cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	// Rules of constants
	// 1. You cannot change constant
	// 2. You cannot initialize a constant to a runtime value
	// 3. You can initialize a constant using built-in function len

	// constant types
	const min = 1 + 1
	const pi = 3.14 * min
	const version = "2.0.1" + "-beta"
	const debug = !true
	fmt.Println(min, pi, version, debug)

	// Multiple declaration
	const minimium, maximium int = 1, 1000
	/*
		const (
			minimium int = 1
			maximium int = 1000
			medium (p.s. this return 1000)
		)
	*/
	fmt.Println(minimium, maximium) // 1, 1000
}
