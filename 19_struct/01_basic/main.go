package main

import "fmt"

func main() {
	/*
		Structs allow us to represent concept -> shape your own type using other type
		Structs are blueprints -> they are `fixed` at compile-time
		It's like `class` in OOP language
		Groups related data in a single type
		You can not add new fields during run-time
	*/

	type person struct {
		name, lastname string
		age            int
	}

	picasso := person{name: "Pable", lastname: "Picasso", age: 91}

	var freud person
	freud.name = "Sigmond"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("\n%s's age is %d\n", picasso.lastname, picasso.age)

	fmt.Printf("\n Person: %#v\n", picasso)
	fmt.Printf("\n Person: %#v\n", freud)
}
