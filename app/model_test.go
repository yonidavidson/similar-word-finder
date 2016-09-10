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

func TestSetGetDB(t *testing.T) {

	db := NewInMemoryDB()
	db.set("key1", "value1")
	db.set("key2", "value2")
	db.set("key3", "value3")
	db.set("key3", "value4")
	s := db.size()
	if s != 4 {
		t.Errorf("db set didn't work - expected:4, result:%d", s)
	}
	valArray, err := db.get("key1")
	if err != nil {
		t.Errorf("failed to get from db key1")
	}
	if !Contains(valArray, "value1") {
		t.Errorf("db.get got wrong values, expected:[value1], result:%v", valArray)
	}

	valArray, err = db.get("key3")
	if err != nil {
		t.Errorf("failed to get from db key3 ")
	}
	if !Contains(valArray, "value3") || !Contains(valArray, "value4") {
		t.Errorf("db.get got wrong values, expected:[value3, value4], result:%v", valArray)
	}
}
