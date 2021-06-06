package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// Print the rune into byte -> decoding
	text := `Go est un langage de programmation compilé et concurrent inspiré de C et Pascal. 
Ce langage a été développé par Google6 à partir d'un concept initial de Robert Griesemer, 
Rob Pike et Ken Thompson. Go possède deux implémentations, la première utilise gc, 
le compilateur Go et la seconde utilise gccgo, « frontend » GCC écrit en C++. 
Go est écrit en C en utilisant yacc et GNU Bison pour l'analyse syntaxique7 jusqu'à la version 1.4, 
et en Go lui-même pour les versions suivantes (1.5).`

	// for i := 0; i < len(text); {
	// 	r, size := utf8.DecodeRuneInString(text[i:])
	// 	fmt.Printf("%c", r)

	// 	i += size
	// }

	for _, r := range text {
		fmt.Printf("%c", r)
	}
	fmt.Println()

	word := []byte("développé")
	fmt.Printf("%s = % [1]x\n", word)

	// var size int
	// for i := range string(word) {
	// 	if i > 0 {
	// 		size = i
	// 		break
	// 	}
	// }

	_, size := utf8.DecodeRune(word)

	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)

}
