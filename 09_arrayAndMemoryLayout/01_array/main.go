package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golder"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"
	// books[4] = "Neverland" // error! out of bounds

	fmt.Printf("books   : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)
	wBooks[0] = books[0]
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	// for i := 0; i < len(sBooks); i++ {
	// 	sBooks[i] = books[i+1]
	// }

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	for _, v := range sBooks {
		fmt.Println(v)
		v += "won't effect"
	}
	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\n Published Books: ")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
