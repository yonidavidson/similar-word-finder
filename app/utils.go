package main

import (
	"errors"
	"sort"
)

type ByAlphabeticalOrder []byte

func (a ByAlphabeticalOrder) Len() int           { return len(a) }
func (a ByAlphabeticalOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAlphabeticalOrder) Less(i, j int) bool { return a[i] < a[j] }

func SortAlphabeticalOrder(s string) string {
	a := []byte(s)
	sort.Sort(ByAlphabeticalOrder(a))
	return string(a)
}

func LoadDataToDb(d Db, path string) (int, error) {
	return -1, errors.New("failed to load data to db")
}
