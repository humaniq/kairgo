package kairgo

import (
	"encoding/json"
	"fmt"
)

// ResponseGallery ...
type ResponseGallery struct {
	RawResponse []byte
	Errors      []Error  `json:"Errors"`
	Status      string   `json:"status"`
	SubjectIDs  []string `json:"subject_ids"`
}

// ViewGallery ...
func (k *Kairos) ViewGallery(galleryName string) (*ResponseGallery, error) {
	if galleryName == "" {
		return nil, fmt.Errorf("galleryName: should be present")
	}

	p := make(map[string]interface{})
	p["gallery_name"] = galleryName

	b, mErr := json.Marshal(p)
	if mErr != nil {
		return nil, mErr
	}

	req, reqErr := k.newRequest("POST", "gallery/view", b)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, doErr := k.do(req)
	if doErr != nil {
		return nil, doErr
	}

	re := &ResponseGallery{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
