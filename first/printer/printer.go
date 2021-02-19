package printer

import "fmt"

// Hello is an exorted function
func Hello() {
	fmt.Println("Unexported hello")
}
