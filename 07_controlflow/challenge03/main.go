package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetToConvert]
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}
	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		// There is an error
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
