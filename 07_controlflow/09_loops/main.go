package main

import "fmt"

func main() {
	// There is no while loop in go, only for
	var (
		sum int
		i   = 1
	)

	// for i := 1; i <= 1000; i++ {
	// 	sum += i
	// }

	// fmt.Println(sum)
	// break a loop
	for {
		if i > 5 {
			break
		}
		sum += i
		i++
	}

	fmt.Println(sum)

	// continue a loop
	sum = 0
	i = 1

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}

		sum += i
		i++
	}
	fmt.Println(sum)
}
