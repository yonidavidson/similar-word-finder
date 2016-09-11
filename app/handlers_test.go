package main

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	verb    string
	url     string
	handler func(c web.C, w http.ResponseWriter, r *http.Request)
}

var TEST_CASES = [...]TestCase{
	{"GET", "/api/v1/similar", SimilarHandler},
	{"GET", "/api/v1/stats", StatsHandler},
}

func TestEndpointsReturnValue(t *testing.T) {
	db = NewInMemoryDB()

	for _, tcase := range TEST_CASES {
		t.Log(tcase)
		req, err := http.NewRequest(tcase.verb, tcase.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := web.HandlerFunc(tcase.handler)
		ctx := web.C{}

		handler.ServeHTTPC(ctx, rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

}

func TestSimilarHandler(t *testing.T) {
	db = NewInMemoryDB()
	_, err := LoadDataToDb(db, "words_clean_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/api/v1/similar?word=stressed", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := web.HandlerFunc(SimilarHandler)
	ctx := web.C{}

	handler.ServeHTTPC(ctx, rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	data := new(SimilarResponse)
	json.Unmarshal(rr.Body.Bytes(), &data)
	words := data.Similar

	if len(words) != 1 {
		t.Fatal("not the number of words in dictionary. expected:%d, result:%d", 1, len(words))
	}

	if Contains(words, "stressed") {
		t.Fatal("should not return queried word , result:%v", words)
	}

	if !Contains(words, "desserts") {
		t.Fatal("didn't find  similar word , result:%v", words)
	}
}

func TestStatsHandler(t *testing.T) {
	db = NewInMemoryDB()
	loadSize, err := LoadDataToDb(db, "words_clean_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	appProps.TotalWords = loadSize

	req, err := http.NewRequest("GET", "/api/v1/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := web.HandlerFunc(StatsHandler)
	ctx := web.C{}

	handler.ServeHTTPC(ctx, rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	data := new(Props)
	json.Unmarshal(rr.Body.Bytes(), &data)
	size := data.TotalWords

	if size != appProps.TotalWords {
		t.Fatal("size of db not matching. expected:%d, result:%d", size, appProps.TotalWords)
	}
}
