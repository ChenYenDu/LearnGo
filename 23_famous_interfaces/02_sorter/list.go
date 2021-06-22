package main

import (
	"sort"
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery.\n"
	}
	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("(%-10T) ----> ", it)
		// it.print()
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {

	for _, p := range l {

		p.discount(ratio)

	}
}

// Len is the number of elements in the collection.
func (l list) Len() int { return len(l) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (l list) Less(i, j int) bool {
	return l[i].title < l[j].title
}

// Swap swaps the elements with indexes i and j.
func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byRelease struct {
	list
}

func (br byRelease) Less(i, j int) bool {
	return br.list[i].released.Before(br.list[j].released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}
