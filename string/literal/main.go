package main

import "fmt"

func main() {
	var s string
	s = "how are you?" // String literal, type is string
	s = `how are you?` // Raw String Literal, type is string
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)
	s = `
	<html>
		<body>"Hello"</body>
	</html>
	`
	fmt.Println(s)
}
