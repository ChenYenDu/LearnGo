package main

import "fmt"

func main() {
	/*
		Inheritance: (do not exists in go)
			`is-a` relationship: book is-a text  -- article is-a text

		for GO, we use Embedding (compoition)
		Embedding:
			`has-a` relationship: book has text -- article has text

	*/

	type text struct {
		title string
		words int
	}

	type book struct {
		// text text
		text  // anonymous field, use the type name automatically
		isbn  string
		title string // parent type has the priority
	}

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	fmt.Printf("%s has %d words (isbn: %s)\n", moby.text.title, moby.text.words, moby.isbn)
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

	fmt.Printf("%#v\n", moby)
}
