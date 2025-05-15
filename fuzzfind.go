package main

import (
	"fmt"

	levenshtein "github.com/ka-weihe/fast-levenshtein"
)

func main() {
	s1 := "fast"
	s2 := "fastest"
	distance := levenshtein.Distance(s1, s2)
	fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
	// => The distance between fast and fastest is 3.
}
