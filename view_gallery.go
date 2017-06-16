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

	p := map[string]interface{}{
		"gallery_name": galleryName,
	}

	resp, err := k.makeRequest("POST", "gallery/view", p)
	if err != nil {
		return nil, err
	}

	re := &ResponseGallery{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
