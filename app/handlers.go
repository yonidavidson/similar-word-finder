package main

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func SimilarHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	payload := new(SimilarResponse)
	word := r.URL.Query().Get("word")
	key := SortAlphabeticalOrder(word)
	data := db.get(key)
	payload.Similar = Exclude(data, word)
	json, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(json))
}

func StatsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(appProps)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(json))
}
