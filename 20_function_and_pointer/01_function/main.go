package main

import "fmt"

func main() {
	fmt.Println(f1("This is p1", 1)) /* 1 Thi is P2 */
	fmt.Println(f2("This is p1", 2)) /* 2 This is P1 */
	squareFunc := returnSquares()
	f := squareFunc()
	fmt.Println(f())
	fmt.Println(f())

	r := returnSquares()
	rTest := r()
	fmt.Println(rTest())
	fmt.Println(rTest())
}

func f1(p1 string, p2 int) (r1 int, r2 string) {
	r1 = p2
	r2 = p1
	return r1, r2
}

func f2(p1 string, p2 int) (r1 int, r2 string) {
	/*
		var r1 int
		var r2 int
	*/
	r1 = p2
	r2 = p1
	return
}

// func sum(vals ...int, strs ...string) int {
// 	total := 0
// 	for _, val := range vals {
// 		total += val
// 	}
// 	for _, s := range strs {
// 		fmt.Println(s)
// 	}
// 	return total
// }

func squares() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

// f -> fun() int{x++ return x * x}

func returnSquares() func() func() int {
	return squares
}
