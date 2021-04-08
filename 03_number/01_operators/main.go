package main

import "fmt"

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	// average = myAge + yourAge (type error! average is float64!)
	average = float64(myAge + yourAge)
	fmt.Println(average)

	// Be careful, float could be inaccuracy!!!
	ratio := 1.0 / 10
	fmt.Printf("%.60f\n", ratio)

	// result of 3/2
	fmt.Println("sum: ", 3+2)    // sum: int
	fmt.Println("sum: ", 2+3.5)  // sum: float64
	fmt.Println("dif: ", 3-1)    // difference: int
	fmt.Println("dif: ", 3-0.5)  // difference: float64
	fmt.Println("prod: ", 4*5)   // product: int
	fmt.Println("prod: ", 5*2.5) // product: float64
	fmt.Println("quot: ", 8/2)   // quotient: int
	fmt.Println("quot: ", 8/1.5) // quotient: float64

	fmt.Println("rem: ", 8%3) // remainder: only for integer
	// fmt.Println("rem: ", 8.0 % 3) // error
	f := 8.0
	fmt.Println("rem: ", int(f)%3)

	// var ratio float64 = 3 / 2 (error!)
	var ratio2 = float64(int(3) / int(2))
	fmt.Println(ratio2)
	fmt.Printf("%T\n", ratio2)
	ratio2 = 3.0 / 2
	fmt.Println(ratio2)

}
