package main

import (
	"fmt"
	"time"
)

func main() {
	type numbers = [5]string
	zero := numbers{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := numbers{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := numbers{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := numbers{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := numbers{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := numbers{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := numbers{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := numbers{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := numbers{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := numbers{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := numbers{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]numbers{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {

		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		curr := time.Now()
		hour, min, sec := curr.Hour(), curr.Minute(), curr.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]numbers{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				// colon blink
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)

	}

}
