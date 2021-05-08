package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"unsafe"
)

const size = 1e7

func main() {

	// Exercise 01: Fix the backing array problems
	nums := []int{56, 89, 15, 25, 30, 50}

	// mine := nums
	mine := append([]int(nil), nums[:3]...)

	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine          :", mine)
	fmt.Println("Original nums :", nums[:3])

	// Exercise 02: Sort the backing array
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	mid := len(items) / 2
	smid := items[mid-1 : mid+2]

	// sort the smid will affect the items
	// as well. their backing array is the same.
	sort.Strings(smid)

	fmt.Println()
	fmt.Println("Sorted  :", items)

	// Exercise 03:

	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var array [size]int

	// 2. print the memory usage (use the report func).
	report("after declaring an array")

	// 3. copy the array to a new array.
	array2 := array

	// 4. print the memory usage
	report("after copying the array")

	// 5. pass the array to the passArray function
	passArray(array)

	// 6. convert one of the arrays to a slice
	slice1 := array[:]

	// 7. slice only the first 1000 elements of the array
	slice2 := array[1e3:]

	// 8. slice only the elements of the array between 1000 and 10000
	slice3 := array[1e3:1e4]

	// 9. print the memory usage (report func)
	report("after slicing")

	// 10. pass the one of the slices to the passSlice function
	passSlice(slice3)
	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size : %d bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size : %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size : %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size : %d bytes.\n", unsafe.Sizeof(slice3))

	// Exercise 04: Obserce length and capacity
	// games := []string{"pacman", "teries"}
	// var games []string
	// fmt.Printf("length: %d,  capacity: %d\n", len(games), cap(games))

	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("length: %d,  capacity: %d\n", len(games), cap(games))
	games := []string{"pacman", "mario", "tetris", "doom"}

	fmt.Println()
	for i := 0; i <= len(games); i++ {
		ele := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(ele), cap(ele))
	}

	fmt.Println()
	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's     len: %d cap: %d\n", len(zero), cap(zero))
	}

	fmt.Println()
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s len: %d cap: %d\n", n, len(s), cap(s))
	}

	fmt.Println()
	zero = zero[:cap(zero)]

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	fmt.Println()
	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  :%q\n", zero)
	fmt.Printf("games :%q\n", games)

	// Exercise 04: Obserce the capacity growth
	var (
		mnums  []int
		oldCap float64
	)

	// loop 10 million times
	for len(mnums) < 10e6 {
		// get the capacity
		c := float64(cap(mnums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(mnums), c, c/oldCap,
			)
		}

		// keep track of the previous capcaity
		oldCap = c

		// append an arbitrary element to the slice
		mnums = append(mnums, 1)
	}

	// Exercise 05 : Correct the lyric
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	lyric = append([]string{"yesterday"}, lyric...)

	const N, M = 8, 13

	lyric = append(lyric, lyric[N:M]...)

	lyric = append(lyric[:N], lyric[M:]...)

	fmt.Printf("%s\n", lyric)

}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
