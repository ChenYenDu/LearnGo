package main

import "fmt"

// Pointer: stores the memory address of a value
// & -> find the address | * -> finds the value | *Type -> donates a pointer type

func main() {
	var (
		counter int
		V       int
		P       *int // declare a pointer to an int
	)

	counter = 100
	P = &counter
	V = *P

	fmt.Printf("counter : %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P       : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V       : %-13d addr: %-13p\n", V, &V)

	fmt.Println("... change counter")
	counter = 10
	fmt.Printf("counter : %-13d addr: %-13p\n", counter, &counter)

	fmt.Println("... change counter in passVal()")
	counter = passVal(counter)
	fmt.Printf("counter : %-13d addr: %-13p\n", counter, &counter)

	fmt.Println("... change counter in passPtrVal()")
	passPtrVal(&counter)
	passPtrVal(&counter)
	passPtrVal(&counter)
	passPtrVal(&counter)
	passPtrVal(&counter)
	fmt.Printf("counter : %-13d addr: %-13p\n", counter, &counter)

	// if P == nil {
	// 	fmt.Printf("P is nil and it's address is %p\n", P) // memory address are represent in hex(16進位) number
	// }

	// if P == &counter {
	// 	fmt.Printf("P is equal to counter's address: %p == %p\n", P, &counter)
	// }
}

func passPtrVal(pn *int) {
	fmt.Printf("pn      : %-13p addr: %-13p *P: %-13d\n", pn, &pn, *pn)

	*pn++ // same as (*pn)++

	fmt.Printf("pn      : %-13p addr: %-13p *P: %-13d\n", pn, &pn, *pn)

	pn = nil // this do not effect the original pointer because it's a copy
}

func passVal(n int) int {
	n = 50
	fmt.Printf("n       : %-13d addr: %-13p\n", n, &n)
	return n
}
