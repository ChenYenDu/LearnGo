package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		Bufio.Scanner:
			Scans an input stream line by line into a buffer (by default)
	*/

	os.Stdin.Close()

	in := bufio.NewScanner(os.Stdin)
	// os.Stdin: Standard input -> very simply put: help us read input that comes from command lines

	// wait until user write something in console
	// but this only save the last input
	// in.Scan()
	// in.Scan()
	// in.Scan() // only save this
	// fmt.Println("Scanned text:", in.Text())

	var lines int

	for in.Scan() {
		// fmt.Println("Scanned text:", in.Text())
		lines++
	}

	fmt.Printf("There are %d lines here.\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	// fmt.Println("Scanned bytes:", in.Bytes())

}
