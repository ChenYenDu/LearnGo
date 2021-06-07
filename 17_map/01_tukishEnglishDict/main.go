package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [chinese word]")
		return
	}

	query := args[0]
	// dict := map[string]string{}
	// dict["up"] = "上"
	// dict["down"] = "下"
	// dict["left"] = "左"
	// dict["right"] = "右"

	dict := map[string]string{
		"up":    "尚",
		"down":  "下",
		"left":  "左",
		"right": "右",
	}

	// fix the wrong translation
	dict["up"] = "上"

	// add an mistake empty string pair
	dict["mistake"] = ""

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)

	fmt.Printf("# of keys: %d\n", len(dict))
	// fmt.Printf("Zero value: %#v\n", dict)

}
