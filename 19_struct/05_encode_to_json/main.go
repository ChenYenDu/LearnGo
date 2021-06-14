package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"` // this is a field settings make it not shown in json but other packages
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"paul", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "666", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(string(out)) // json only encodes the exported fields
}
