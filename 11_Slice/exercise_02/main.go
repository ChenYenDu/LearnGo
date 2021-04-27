package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	// append exercise - 1
	png, headers := []byte{'P', 'N', 'G'}, []byte{}

	headers = append(headers, png...)

	isEqual := bytes.Equal(png, headers)

	fmt.Println("png is equal to header?", isEqual)

	// append exercise - 2
	var pizza []string
	var departure []time.Time
	var graduations []int
	var isOn []bool

	pizza = append(pizza, "pepperoni", "extra", "cheese")
	graduations = append(graduations, 1998, 2005, 2018)
	now := time.Now()
	departure = append(departure, now, now.Add(time.Hour*24))
	isOn = append(isOn, true, false)

	fmt.Println(pizza, graduations, departure, isOn)

	// append exercise - 3 (fix bug)'// toppings := []int{"black olives", "green peppers"}

	toppings := []string{"black olives", "green peppers"}
	var pizz []string
	pizz = append(pizz, toppings...)
	pizz = append(pizz, "onions")
	pizz = append(pizz, "extra", "cheese")

	fmt.Printf("pizza       : %s\n", pizz)

	// append exercise - 4 (append and sort)
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	var nums []int
	for _, s := range args {
		n, err := strconv.Atoi(s)

		if err != nil {
			continue
		}

		nums = append(nums, n)
	}

	sort.Ints(nums)
	fmt.Println(nums)

	// append exercise - 4 (house price)
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")
	for _, row := range rows {
		cols := strings.Split(row, separator)
		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
		fmt.Println()
	}

	var avg []float64

	var (
		sizeAvg  float64
		bedAvg   float64
		bathAvg  float64
		priceAvg float64
	)

	for i, _ := range locs {
		sizeAvg += float64(sizes[i])
		bedAvg += float64(beds[i])
		bathAvg += float64(baths[i])
		priceAvg += float64(prices[i])
	}

	l := float64(len(locs))
	sizeAvg /= l
	bedAvg /= l
	bathAvg /= l
	priceAvg /= l

	avg = append(avg, sizeAvg, bedAvg, bathAvg, priceAvg)

	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	fmt.Printf("%-15s", "")
	for _, v := range avg {
		fmt.Printf("%-15v", v)
	}
	fmt.Println()

}
