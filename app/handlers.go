package main

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	data, err := getUserData(c.URLParams["username"])
	if err != nil {
		fmt.Fprintf(w, "Failed to get data:", 500)
	} else {
		json, _ := json.Marshal(data)
		fmt.Fprintf(w, "%v", string(json))
	}
}
