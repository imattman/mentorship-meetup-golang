package main

import "fmt"

func main() {
	scoreFn := buildScorer(runeScores())

	for _, word := range loadWords() {
		s := scoreFn(word)
		fmt.Printf("%s : %d\n", word, s)
	}
}

func loadWords() []string {
	words := []string{
		"apple",
		"orange",
		"banana",
		"pineapple",
		"cherry",
	}

	return words
}

func runeScores() map[rune]int {
	runeVals := map[rune]int{}

	return runeVals
}

func buildScorer(rVals map[rune]int) func(string) int {
	fn := func(word string) int {
		total := 0
		for _, r := range []rune(word) {
			runeVal := rVals[r]
			total += runeVal
		}

		return total
	}

	return fn
}
