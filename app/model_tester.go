package main

import (
	"testing"
)

func TestDbEmptyInit(t *testing.T) {

	db := new(InMemoryDB)
	s := db.size()
	if s > 0 {
		t.Errorf("db not empty on init - expected:0, result:" + string(s))
	}

}
