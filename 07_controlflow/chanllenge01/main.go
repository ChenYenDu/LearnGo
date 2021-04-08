package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
Usage: [username] [password]
`

func main() {
	if len(os.Args) != 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	username, password := os.Args[1], os.Args[2]

	if username != "jack" {
		fmt.Printf("Access denied for %s.\n", username)
		return
	} else {
		if password != "1888" {
			fmt.Printf("Invalid password for \"jack\". \n")
			return
		} else {
			fmt.Printf("Access granted to \"jack\". \n")
			return
		}
	}

}
