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

func TestLoadDb(t *testing.T) {

	db := new(InMemoryDB)
	size, err := db.load("words_clean_sample.txt")
	if err != nil {
		t.Errorf("failed to load to DB")
	}
	if size <= 0 {
		t.Errorf("failed to load to DB expected: >0, result:%d", size)
	}
}
