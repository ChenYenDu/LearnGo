package main

import (
	"fmt"
	"strconv"
)

func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("show -> local = %d\n", local)

	local = incr(local)
	fmt.Printf("show -> local = %d\n", local)

	local = incrBy(local, 5)
	fmt.Printf("show -> local = %d\n", local)

	_, err := incrByStr(local, "TWO")

	if err != nil {
		fmt.Printf("err -> %d\n", err)
	}

	local, _ = incrByStr(local, "2")
	fmt.Printf("show -> local = %d\n", local)

	show(incrBy(local, 2))
	show(local)

	local = santize(incrByStr(local, "2"))
	show(local)

	local = limit(incrBy(local, 5), 500)
	show(local)

}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n * factor
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)

	n = incrBy(n, m)
	return n, err
}

func santize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

func limit(n int, lim int) (m int) {
	m = n
	if m >= lim {
		return lim
	}
	return
	// when there is named return variable, this equals to `return m`
}
