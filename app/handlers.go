package main

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

func SimilarHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	payload := new(SimilarResponse)
	key := SortAlphabeticalOrder(r.URL.Query().Get("word"))
	data := db.get(key)
	payload.Similar = data
	json, _ := json.Marshal(payload)
	log.Printf("***********%s", string(json))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "%v", string(json))
}

func StatsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "StatsHandler: not ready")
}
