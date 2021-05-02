package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 01 slicing the number
	data := "2 4 6 1 3 5"

	var nums []int
	var evens []int
	var odds []int

	for _, v := range strings.Fields(data) {
		num, err := strconv.Atoi(v)

		if err != nil {
			continue
		}

		nums = append(nums, num)

		if (num % 2) == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}

	}

	l := len(nums)

	fmt.Printf("nums         : %v\n", nums)
	fmt.Printf("evens        : %v\n", evens)
	fmt.Printf("odds         : %v\n", odds)
	fmt.Printf("middle       : %v\n", nums[l/2:l/2+2])
	fmt.Printf("first 2      : %v\n", nums[:2])
	fmt.Printf("last 2       : %v\n", nums[l-2:])
	fmt.Printf("evens last 1 : %v\n", evens[len(evens)-1:])
	fmt.Printf("odds last 2  : %v\n", odds[len(odds)-2:])
	fmt.Println()

	// 02 slice with arguments
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("%q\n\n", ships)

	from, to := 0, len(ships)

	switch poss := os.Args[1:]; len(poss) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 2:
		to, _ = strconv.Atoi(poss[1])
		fallthrough
	case 1:
		from, _ = strconv.Atoi(poss[0])
	}

	if l := len(ships); from < 0 || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])

	// 03 slice housing prices
	const (
		datas = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// parse the data
	rows := strings.Split(datas, "\n")
	cols := strings.Split(rows[0], separator)

	// default case: slice for all the columns
	from, to = 0, len(cols)

	// find the from and to positions depending on the command-line arguments
	args := os.Args[1:]
	for i, v := range cols {
		l := len(args)

		if l >= 1 && v == args[0] {
			from = i
		}

		if l == 2 && v == args[1] {
			to = i + 1 // +1 because the stopping index is a position
		}
	}

	// from cannot be greater than to: reset invalid arg to 0
	if from > to {
		from = 0
	}

	for i, row := range rows {
		cols := strings.Split(row, separator)

		// print only the requested columns
		for _, h := range cols[from:to] {
			fmt.Printf("%-15s", h)
		}

		fmt.Println()

		// print extra new line for the header
		if i == 0 {
			fmt.Println()
		}
	}

}
