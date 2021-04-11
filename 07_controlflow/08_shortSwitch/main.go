package main

func main() {
	switch i := 10; { // switch by default true
	case i > 0:
		// positive
	case i < 0:
		// negative
	default:
		// zero
	}
}
