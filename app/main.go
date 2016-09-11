package main

import (
	"github.com/zenazn/goji"
	"log"
)

var db Db
var appProps Props

func main() {
	db = NewInMemoryDB()
	size, err := LoadDataToDb(db, "./words_clean.txt")
	if err != nil {
		log.Fatal("failed to load DB. error=" + err.Error())
	}
	log.Println("loaded %d values to db", size)
	goji.Get("/api/v1/similar", SimilarHandler)
	goji.Get("/api/v1/stats", StatsHandler)
	goji.Serve()
}
