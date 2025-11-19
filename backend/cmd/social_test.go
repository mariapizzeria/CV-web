package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSocialSuccess(t *testing.T) {
	tester := httptest.NewServer(App())
	defer tester.Close()

	res, err := http.Get(tester.URL + "/social")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Status code is %d", res.StatusCode)
	}
}
