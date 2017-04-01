package main

import (
	"errors"
	"fmt"
	"strconv"
)

// START MAIN OMIT
func main() {
	for i := 1; i < 25; i++ {
		// can ignore the error using '_' place holder
		// variable assignment needs to match number of values returned
		result, _ := fizzbuzz(i)

		fmt.Printf("%2d: %s\n", i, result)
	}
}

// END MAIN OMIT

// START FIZZBUZZ OMIT
func fizzbuzz(n int) (string, error) {
	if n > 20 {
		return "", errors.New("Sorry, I don't FizzBuzz that high")
	}

	// switch/case - no fall through by default
	// requires explicit use of 'fallthrough' keyword
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz!", nil
	case n%3 == 0:
		return "Fizz!", nil
	case n%5 == 0:
		return "Buzz!", nil
	default:
		return strconv.Itoa(n), nil
	}
}

// END FIZZBUZZ OMIT
