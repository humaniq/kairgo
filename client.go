package kairgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	pkgVersion     = "0.0.1"
	defaultBaseURL = "https://api.kairos.com/"
	userAgent      = "kairgo/" + pkgVersion
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Kairos ...
type Kairos struct {
	// HTTP client
	client httpClient

	// Base URL for API request
	baseURL *url.URL

	// User agent used when communicating with the Kairos API
	UserAgent string

	credentials struct {
		appID  string
		appKey string
	}
}

// New returns a new Kairos API client.
func New(api, appID, appKey string, hc httpClient) (*Kairos, error) {
	if api == "" {
		api = defaultBaseURL
	}

	if !strings.HasSuffix(api, "/") {
		api += "/"
	}

	bURL, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	if hc == nil {
		hc = http.DefaultClient
	}

	k := &Kairos{}
	k.client = hc
	k.baseURL = bURL
	k.UserAgent = userAgent
	k.credentials.appID = appID
	k.credentials.appKey = appKey

	return k, nil
}

func (k *Kairos) newRequest(method, path string, body []byte) (*http.Request, error) {
	rel, pErr := url.Parse(path)
	if pErr != nil {
		return nil, pErr
	}

	uri := k.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set up header
	req.Header.Add("User-Agent", k.UserAgent)
	req.Header.Add("app_id", k.credentials.appID)
	req.Header.Add("app_key", k.credentials.appKey)

	return req, nil
}

func (k *Kairos) do(req *http.Request) ([]byte, error) {
	resp, doErr := k.client.Do(req)
	if doErr != nil {
		return nil, doErr
	}

	defer resp.Body.Close()

	b, rErr := ioutil.ReadAll(resp.Body)
	if rErr != nil {
		return nil, rErr
	}

	if resp.StatusCode != http.StatusOK {
		var e interface{}
		uErr := json.Unmarshal(b, &e)
		if uErr != nil {
			return nil, uErr
		}
		return nil, fmt.Errorf("Kairos returned a status of non 200 %+v", e)
	}

	return b, nil
}
