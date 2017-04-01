package main

import "fmt"

func fib(n int) int {
	a, b := 0, 1 // HL

	for i := 0; i < n; i++ {
		a, b = b, a+b // HL
	}

	return a
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, fib(i))
	}
}
