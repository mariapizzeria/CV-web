package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEducationSuccess(t *testing.T) {
	tester := httptest.NewServer(App())
	defer tester.Close()

	res, err := http.Get(tester.URL + "/education")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Status code is %d", res.StatusCode)
	}
}
