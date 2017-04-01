package main

import (
	"fmt"
	"unicode"
)

var shift = 17

// START MAIN OMIT
var message = `Kn
bdan
cx
marwt
hxda
Xejucrwn`

func main() {
	// the character type is known as a 'rune' in Go
	// runes are UTF-8 code points (1-4 bytes)
	clear := []rune(message) // convert string to slice of runes
	encoded := []rune{}      // empty slice of encoded runes

	for _, c := range clear {
		if unicode.IsLetter(c) {
			c = shiftN(c, 17) // shift 17
		}
		// build up the encoded slice one rune at a time
		encoded = append(encoded, c)
	}

	// encoded runes converted back to string for output
	fmt.Println(string(encoded))
}

// END MAIN OMIT

// START SHIFT OMIT
func shiftN(c rune, n int) rune {
	// Go doesn't allow the ambiguity of *implicit* type conversions for interop
	// between different types (ints, floats, runes, etc.)
	//
	// In this func runes are explicity converted to ints for interop then back
	// to a rune for the return value

	asciiOffset := int('a') // rune conversion to int
	if unicode.IsUpper(c) {
		asciiOffset = int('A') // rune conversion to int
	}

	charIdx := int(c) - asciiOffset
	shifted := ((charIdx + n) % 26) + asciiOffset

	return rune(shifted) // int converted to rune
}

// END SHIFT OMIT
