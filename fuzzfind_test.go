package fuzzfind_test

import (
	"fmt"
	"fuzzfind"
	"testing"
)

func TestFuzzfindTwoStrings(t *testing.T) {
	fmt.Print("Testing fuzzyfind...\n")
	t.Parallel()
	var want int = 3
	got := fuzzfind.FuzzfindTwoStrings("fast", "fastest")
	if want != got {
		t.Errorf("FuzzfindTwoStrings() = %d, want %d", got, want)
	}
}
