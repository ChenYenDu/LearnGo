package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"

	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
	}

	args := os.Args

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass ||
		u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
