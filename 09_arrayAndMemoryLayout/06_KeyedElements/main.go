package main

import "fmt"

func main() {
	rates01 := [3]float64{
		2: 1.5,
	}

	fmt.Printf("rates: %v\n", rates01) // [0, 0, 1.5]

	rates02 := [...]float64{
		5: 1.5,
		2.5,
		0: 0.5,
	}

	fmt.Printf("rates: %v\n", rates02) //[0.5, 0, 0, 0 ,0, 1.5, 2.5]

	const (
		ETH = 9 - iota // 9th
		WAN            // 8th
		ICX            // 7th
	)
	rates := [...]float64{
		WAN: 120.5, // wanchain
		ETH: 25.5,  // ethereum
		ICX: 20,
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}
