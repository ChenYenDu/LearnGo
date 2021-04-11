package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.Now().Hour()
	// fmt.Printf("It's %d o'clock.\n", t)

	switch t := time.Now().Hour(); {
	case t >= 18:
		fmt.Println("good evening")
	case t >= 12:
		fmt.Println("good afternoon")
	case t >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good morning")
	}

}
