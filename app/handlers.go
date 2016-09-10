package main

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func SimilarHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	payload := new(SimilarResponse)
	key := SortAlphabeticalOrder(r.URL.Query().Get("word"))
	data := db.get(key)
	payload.similar = data
	json, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(json))
}

func StatsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "StatsHandler: not ready")
}
