package main

import (
	"fmt"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please type a Hogwart house name")
		return
	}

	hogwarts := make(map[string][]string, 4)
	hogwarts["gryffindor"] = []string{
		"weasley", "hagrid", "dumbledore", "lupin",
	}
	hogwarts["hufflepuf"] = []string{
		"wenlock", "scamander", "helga", "diggory",
	}
	hogwarts["ravenclaw"] = []string{
		"flitwick", "bagnold", "wildsmith", "montmorency",
	}
	hogwarts["slytherin"] = []string{
		"horace", "nigellus", "higgs", "scorpius", "wizardry", "unwanted",
	}
	hogwarts["bobo"] = []string{
		"wizardry", "unknown",
	}

	// delete bobo
	delete(hogwarts, "bobo")

	house, students := args[0], hogwarts[args[0]]

	if students == nil {
		fmt.Printf("Sorry, I know nothing about %s\n", house)
		return
	}

	// only sort the clone
	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)

	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}

	fmt.Println()

	// House        Student Name
	// ---------------------------
	// gryffindor   weasley
	// gryffindor   hagrid
	// gryffindor   dumbledore
	// gryffindor   lupin
	// hufflepuf    wenlock
	// hufflepuf    scamander
	// hufflepuf    helga
	// hufflepuf    diggory
	// ravenclaw    flitwick
	// ravenclaw    bagnold
	// ravenclaw    wildsmith
	// ravenclaw    montmorency
	// slytherin    horace
	// slytherin    nigellus
	// slytherin    higgs
	// slytherin    scorpius
	// bobo         wizardry
	// bobo         unwanted
}
