package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "carl"
	fmt.Println(
		len(name),
	)

	name = "杜俊彥"
	fmt.Println(len(name)) // return bit length

	fmt.Println(utf8.RuneCountInString(name)) // return length of utf string
}
