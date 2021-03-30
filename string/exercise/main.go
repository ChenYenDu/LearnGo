package main

import (
	"fmt"
	"os"
	"strings"
)

// string
func main() {
	msg := os.Args[1]
	l := len(msg)

	s := msg + strings.Repeat("!", l)

	fmt.Println(s)

	s = strings.ToUpper(s)

	fmt.Println(s)

}
