package main

import (
	"fmt"

	"github.com/inancgumus/learngo/09-go-type-system/05-defined-types/03-underlying-types/weights"
)

type gram float64
type ounce float64

func main() {
	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var g gram = 1000
	var o ounce

	o = ounce(g) * 0.035273

	fmt.Printf("%g grmas is %.2f ounce\n", g, o)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)
	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(apples)))

	salt = Gram(weights.Gram(100))
	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)
	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of main.Gram: %T\n", Gram(1))

	type myGram weights.Gram

	var pepper myGram = 20
	pepper = myGram(salt)

	fmt.Printf("Type of pepper   : %T\n", pepper)
}
