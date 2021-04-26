package main

import (
	"fmt"
	"time"
)

func main() {

	for shift := 0; ; shift++ {

		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		curr := time.Now()
		hour, min, sec := curr.Hour(), curr.Minute(), curr.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		// alarmed := sec%10 == 0

		for line := range clock[0] {
			l := len(clock)

			s, e := shift%l, l

			if shift%(l*2) >= l {
				s, e = 0, s
			}

			for j := 0; j < l-e; j++ {
				fmt.Print("   ")
			}

			for i := s; i < e; i++ {
				next := clock[i][line]

				if clock[i] == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, " ")
			}

			// for index, digit := range clock {
			// 	// colon blink
			// 	next := clock[index][line]
			// 	if digit == colon && sec%2 == 0 {
			// 		next = "   "
			// 	}
			// 	fmt.Print(next, " ")
			// }
			fmt.Println()
		}

		time.Sleep(time.Second)

	}

}
