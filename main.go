package main

import "fmt"

func pipe(number int, function ...func(int) int) int {

	for _, f := range function {
		number = f(number)
	}

	return number
}

func main() {

	increase := func(n int) int {
		return n + 1
	}

	square := func(n int) int {
		return n * n
	}

	r := pipe(1, increase, square)

	fmt.Println(r)
}
