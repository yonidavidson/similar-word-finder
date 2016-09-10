package main

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func SimilarHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "SimilarHandler: not ready")
}

func StatsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "StatsHandler: not ready")
}
