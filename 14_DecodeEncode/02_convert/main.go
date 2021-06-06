package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	str := "y¨gen ☯︎ ♞"

	/*
		String value are read-only byte slices
		string value => read-only []byte
	*/

	/*
		Convert string to byte creates a []btype, and copies the bytes of
		the string to the new slice's backing array.

		So, change byte do not change the original str
	*/
	bytes := []byte(str)
	// bytes[0] = 'N'
	// bytes[1] = 'O'

	// str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()
	/*
		range go through rune not bytes
		so, index here could get with space larger than 1
	*/
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte    : %c\n", str[0])
	fmt.Printf("2nd byte    : %c\n", str[1])
	fmt.Printf("2nd rune    : %s\n", str[1:3]) // slice a string also return a string value

	/*
		Runes in a UTF-8 encoded string can have `a different number of bytes`
		because UTF-8 is a variable byte-length encoding

		Especially in scripting languages you can index strings easily,
		However, Go doesn't allow you to do so by default because of efficiency reasons.

		Go never hides the cost of doing something.
	*/

	// An easier way but inefficient one
	runes := []rune(str) // create and copy rune one by one new slice backing array

	fmt.Println()
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d runes\n", len(runes))
	fmt.Printf("1st rune:     %c\n", runes[0])
	fmt.Printf("2nd rune:     %c\n", runes[1])
	fmt.Printf("first five:   %c\n", runes[:5])

}
