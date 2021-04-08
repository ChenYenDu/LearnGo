package main

import (
	"fmt"
	"time"
)

func main() {
	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

	// you can only assign value in with the same type
	// error: speed = "100"

	// var running bool
	// running = 1 return error 1 is number but 'running' is bool

	var force float64
	// speed = force reture error
	force = 1 // 1 is number and 'force' is float64, this works.
	fmt.Println(force)

	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 28
	famous = false
	fmt.Println(name, age, famous)

	var prevName string
	prevName = name
	name = "Einstein"

	fmt.Println("previous name: ", prevName)
	fmt.Println("name: ", name)

	// multiple assignments
	var (
		speed2 int
		now    time.Time
	)

	speed2, now = 100, time.Now()

	fmt.Println(speed2, now)

	// swap variable by using multiple assignments
	var (
		currSpeed = 50
		prevSpeed = 100
	)

	currSpeed, prevSpeed = prevSpeed, currSpeed

	fmt.Println(currSpeed, prevSpeed)

}
