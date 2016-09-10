package main

import (
	"errors"
)

type Db interface {
	size() int
	get(key string) (val string, err error)
	set(key string, val string)
}

type InMemoryDB struct {
	m map[string][]string
}

func (d InMemoryDB) size() int {
	return len(d.m)
}

func (d InMemoryDB) get(key string) (val []string, err error) {
	empty := []string{"empty"}
	return empty, errors.New("failed to get key:" + key)
}

func (d InMemoryDB) set(key string, val string) {

}
