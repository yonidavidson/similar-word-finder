package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

var db Db

func main() {

	goji.Get("/api/v1/similar", SimilarHandler)
	goji.Get("/api/v1/stats", StatsHandler)
	goji.Use(middleware.EnvInit)
	goji.Use(JsonText)
	goji.Serve()
}
