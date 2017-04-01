package main

import "fmt"

func main() {
	// long way to declare variable
	var longWay string
	longWay = "long way"
	fmt.Printf("%T : %s\n", longWay, longWay)

	// shorter way with type inference ':='
	shortWay := "short way"
	fmt.Printf("%T : %s\n\n", shortWay, shortWay)

	// many types in one declaration!
	num, str, truth := 42, "yahoo!", true

	// %v translates to the default representation for a type
	fmt.Printf("%T : %v\n", num, num)
	fmt.Printf("%T : %v\n", str, str)
	fmt.Printf("%T : %v\n", truth, truth)
}
