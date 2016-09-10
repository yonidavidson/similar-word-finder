package main

import (
	"bufio"
	"errors"
	"os"
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
	counter := 0
	file, err := os.Open(path)
	if err != nil {
		return counter, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key := SortAlphabeticalOrder(scanner.Text())
		val := scanner.Text()
		d.set(key, val)
		counter++
	}

	if err := scanner.Err(); err != nil {
		return counter, err
	}

	if counter != d.size() {
		return counter, errors.New("error in size after loading data to db ")
	}

	return counter, nil
}
