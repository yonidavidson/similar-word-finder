package main

import (
	"testing"
)

type sortCase struct {
	before string
	after  string
}

type excludeCase struct {
	before []string
	after  []string
	key    string
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
	if size != db.size() {
		t.Errorf("miss match in db size. expected:%d , result:%d", db.size(), size)
	}

	t.Log("size of loaded db is:%d ", size)
	t.Log("db view:%v", db)
}

func TestExclude(t *testing.T) {
	useCases := []excludeCase{
		{[]string{"abcd"}, []string{}, "abcd"},
		{[]string{"abcd", "dbca"}, []string{"dbca"}, "abcd"},
		{[]string{}, []string{}, "abcd"},
	}
	for _, sc := range useCases {
		key := sc.key
		expected := sc.after
		result := Exclude(sc.before, key)
		if Contains(result, key) {
			t.Errorf("failed to exlude key expected: %v, result: %v, key:%s", expected, result, key)
		}
	}
}
