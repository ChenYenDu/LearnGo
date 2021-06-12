package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("[english words] -> [chinese words]")
		return
	}

	query := args[0]

	dict := map[string]string{
		"good":    "好",
		"great":   "很好",
		"perfect": "完美",
		"awesome": "完美",
	}

	// you don't need to check if the key exists or not before deleting it
	delete(dict, "awesome")
	delete(dict, "awesome")
	delete(dict, "hahaha")

	/*
		chinese := dict
		chinese["上"] = "good" // this operation will change both dict and chinese
	*/

	// you should do this if you don't want to change the original map while manipulating the new map
	chinese := make(map[string]string, len(dict))

	for k, v := range dict {
		chinese[v] = k
	}

	// fmt.Printf("english: %q -> chinese: %q", dict, chinese)
	if value, ok := chinese[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)

	fmt.Printf("# of keys: %d\n", len(dict))
	// fmt.Printf("Zero value: %#v\n", dict)

	// Do this if you want to delete all key-value pairs in map

	for k := range dict {
		delete(dict, k)
	}
}
