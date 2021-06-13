package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	// create a regular expression to find any thing that is not a letter
	rx := regexp.MustCompile(`[^a-z]+`)

	// get query from the user
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please input a word to lookup.")
		return
	}

	query := args[0]

	// Read the input file
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		// get the word in lowercase
		word := strings.ToLower(in.Text())

		// save only the letters
		word = rx.ReplaceAllString(word, "")

		// we only want the word that is more than 3 letters
		if len(word) > 2 {
			words[word] = true
		}
	}

	// fmt.Println("sun: ", words["sun"], "tesla: ", words["tesla"])

	// check if the query word exists in the input file
	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}

	fmt.Printf("Sorry, the input doesn't contains %q.\n", query)

}
