package fuzzfind

import (
	"fmt"

	levenshtein "github.com/ka-weihe/fast-levenshtein"
)

func FuzzfindTwoStrings(word1, word2 string) int {
	distance := levenshtein.Distance(word1, word2)
	fmt.Printf("The distance between '%s' and '%s' is %d.\n", word1, word2, distance)
	return distance
	// => The distance between fast and fastest is 3.
}
