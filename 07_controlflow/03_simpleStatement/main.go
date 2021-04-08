package main

import (
	"fmt"
	"strconv"
)

func main() {
	// simple statement allow you to execute a statement inside another statement
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There was no error, n is", n)
	}
}
