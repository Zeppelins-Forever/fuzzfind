package fuzzfind_test

import (
	"fmt"
	"fuzzfind"
	"testing"
)

func TestFuzzfindTwoStrings(t *testing.T) {
	fmt.Print("Testing fuzzyfind...\n")
	type testCase struct {
		a, b string
		want int
	}
	testCases := []testCase{
		{"According to all known laws of aviation, there is no way a bee should be able to fly.", "According to all known laws of there, aviation is no way a bee should be able to fly.", 14},
		{"fast", "fastest", 3},
		{"", "", 0},
		{"the quick brown fox jumps over the lazy frog", "the quick brown fox jumps over the frog lazy", 8},
	}
	for _, tc := range testCases {
		got := fuzzfind.FuzzfindTwoStrings(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("FuzzfindTwoStrings(%q, %q) = %q, want %q", tc.a, tc.b, got, tc.want)
		}
	}
}
