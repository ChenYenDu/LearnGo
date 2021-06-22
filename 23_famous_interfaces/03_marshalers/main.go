package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `[
	{
	 "Title": "moby dick",
	 "Price": 10,
	 "Released": 118281600
	},
	{
	 "Title": "odyssey",
	 "Price": 15,
	 "Released": 733622400
	},
	{
	 "Title": "hobbit",
	 "Price": 25,
	 "Released": -62135596800
	}
   ]`

func main() {

	var l list
	err := json.Unmarshal([]byte(data), &l)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(l)

	// store := list{
	// 	{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
	// 	{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
	// 	{Title: "hobbit", Price: 25},
	// }

	// data, err := json.MarshalIndent(store, "", " ")

	// if err != nil {
	// 	log.Fatal(err)
	// 	// same as:
	// 	// fmt.Println(err)
	// 	// return
	// }

	// fmt.Println(string(data))

}
