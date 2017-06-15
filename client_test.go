package kairgo_test

import (
	"github.com/humaniq/kairgo"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	APP_ID  = "TEST_APP_ID"
	APP_KEY = "TEST_APP_KEY"
)

var (
	mux *http.ServeMux

	client *kairgo.Kairos

	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	// Test client
	server = httptest.NewServer(mux)
	// With test baseUrl
	url, _ := url.Parse(server.URL)
	client, _ = NewClient(url.String())
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func NewClient(baseUrl string) (*kairgo.Kairos, error) {
	return kairgo.New(baseUrl, APP_ID, APP_KEY, nil)
}
