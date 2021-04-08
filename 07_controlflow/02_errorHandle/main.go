package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// strconv.Itoa never fails
	s := strconv.Itoa(42)
	fmt.Println(s)

	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted number    :", n)
	fmt.Println("Return error value  :", err)
}
