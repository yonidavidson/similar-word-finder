package main

import (
	"github.com/zenazn/goji"
)

var db Db

func main() {

	goji.Get("/api/v1/similar", SimilarHandler)
	goji.Get("/api/v1/stats", StatsHandler)
	goji.Serve()
}
