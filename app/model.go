package main

import (
	"errors"
)

type Db interface {
	size() int
}

type InMemoryDB struct {
	m map[string][]string
}

func (d InMemoryDB) size() int {
	return len(d.m)
}

func (d InMemoryDB) load(filePath string) (size int, err error) {
	return 0, errors.New("InMemoryDB: failed to load")
}
