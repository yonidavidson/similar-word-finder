package main

type Db interface {
	size() int
}

type InMemoryDB struct {
	m map[string][]string
}

func (d InMemoryDB) size() int {
	return len(d.m)
}
