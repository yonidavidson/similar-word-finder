package main

import (
	"fmt"
	"time"
)

type Db interface {
	size() int
	get(key string) (val []string)
	set(key string, val string)
}

type SimilarResponse struct {
	Similar []string
}

type Props struct {
	TotalWords          int
	TotalRequests       int
	AvgProcessingTimeNs int64
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

func updateProps(start time.Time) {
	elapsedNs := time.Since(start).Nanoseconds()
	fmt.Println("%d", elapsedNs)
}

func readProps() Props {
	p := new(Props)
	return *p
}

func propsHandler(numOfWords int, updater chan int64, reader chan chan Props) {
	p := new(Props)
	p.TotalWords = numOfWords
	select {
	case elapsedNs := <-updater:
		p.TotalRequests++
		p.AvgProcessingTimeNs = Average(p.AvgProcessingTimeNs, elapsedNs, p.TotalRequests)
	case c := <-reader:
		c <- *p
	default:
		fmt.Println("no activity")
	}

}
