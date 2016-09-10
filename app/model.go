package main

type Db interface {
	size() int
	get(key string) (val []string)
	set(key string, val string)
}

type SimilarResponse struct {
	Similar []string
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

func (d InMemoryDB) get(key string) []string {
	empty := []string{}
	valArray, ok := d[key]
	if !ok {
		return empty
	}
	return valArray
}

func (d InMemoryDB) set(key string, val string) {
	d[key] = append(d[key], val)
}
