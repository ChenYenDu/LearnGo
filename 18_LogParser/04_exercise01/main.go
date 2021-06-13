package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	rx := regexp.MustCompile("[^A-Za-z]+")

	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	total, words := 0, make(map[string]int)

	for in.Scan() {
		word := rx.ReplaceAllString(in.Text(), "")
		word = strings.ToLower(word)

		total++
		words[word]++
	}
	fmt.Printf("There are %d words, %d of than are unique.\n",
		total, len(words))
}
