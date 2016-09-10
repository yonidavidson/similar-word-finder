package main

import (
	"testing"
)

type sortCase struct {
	before string
	after  string
}

func TestSortWord(t *testing.T) {
	useCases := []sortCase{{"abcd", "abcd"}, {"dcba", "abcd"}}
	for _, sc := range useCases {
		expected := sc.after
		result := SortAlphabeticalOrder(sc.before)
		if expected != result {
			t.Errorf("failed to sort word expected: %s, result: %s", expected, result)
		}
	}
}
