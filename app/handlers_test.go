package main

import (
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
