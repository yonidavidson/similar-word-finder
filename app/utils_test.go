package main

import (
	"testing"
)

type sortCase struct {
	before string
	after  string
}

func TestSortWord(t *testing.T) {
	useCases := []sortCase{{"abcd", "abcd"}, {"dcba", "abcd"}, {"dcbbbab", "abbbbcd"}}
	for _, sc := range useCases {
		expected := sc.after
		result := SortAlphabeticalOrder(sc.before)
		if expected != result {
			t.Errorf("failed to sort word expected: %s, result: %s", expected, result)
		}
	}
}

func TestDBLoad(t *testing.T) {
	db := NewInMemoryDB()
	size, err := LoadDataToDb(db, "words_clean_sample.txt")
	if err != nil {
		t.Errorf("failed to load db. result:%v", db)
	}
	if size < 5 {
		t.Errorf("failed to load all data to db. result:%v", db)
	}
}
