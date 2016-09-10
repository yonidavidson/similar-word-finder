package main

import (
	"errors"
)

type Db interface {
	size() int
	get(key string) (val []string, err error)
	set(key string, val string)
}

type InMemoryDB map[string][]string

func NewInMemoryDB() InMemoryDB {
	return make(map[string][]string)
}

func (d InMemoryDB) size() int {
	s := 0
	for _, v := range d {
		s += len(v)
	}
	return s
}

func (d InMemoryDB) get(key string) (val []string, err error) {
	empty := []string{"empty"}
	valArray, ok := d[key]
	if !ok {
		return empty, errors.New("failed to get key:" + key)
	}
	return valArray, nil
}

func (d InMemoryDB) set(key string, val string) {
	d[key] = append(d[key], val)
}
