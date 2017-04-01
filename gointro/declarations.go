package main

import "fmt"

func main() {
	// long way to declare variable
	var longWay string
	longWay = "long way"

	// also valid
	// var longWay string = "another long way"
	fmt.Printf("%T : %q\n", longWay, longWay)

	// shorter way with type inference
	shortWay := "short way"
	fmt.Printf("%T : %q\n", shortWay, shortWay)
}
