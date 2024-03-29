package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
	Decodin process:
		json -> byte -> user
*/

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms,omitempty"`
}

func main() {

	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	// fmt.Println(string(input))

	var users []user

	// err := json.Unmarshal(input, &users)
	// `&` find the address of variables in memory

	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println("err")
		return
	}

	// fmt.Println(users)

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {

		case p == nil:
			fmt.Print(" This user has no power.! ")
		case p["admin"]:
			fmt.Print(" it's an admin ")
		case p["write"]:
			fmt.Print(" can write. ")

		}

		fmt.Println()
	}

}
