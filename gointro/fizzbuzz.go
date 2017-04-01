package main

import (
	"errors"
	"fmt"
	"strconv"
)

// START MAIN OMIT
func main() {
	for i := 1; i < 100; i++ {
		result, err := fizzbuzz(i)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			break
		}

		fmt.Printf("%2d: %s\n", i, result)
	}
}

// END MAIN OMIT

// START FIZZBUZZ OMIT
func fizzbuzz(n int) (string, error) {
	if n > 20 {
		return "", errors.New("Sorry, I don't FizzBuzz that high")
	}

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
