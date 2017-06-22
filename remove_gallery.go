package kairgo

import (
	"encoding/json"
	"fmt"
)

type ResponseRemoveGallery struct {
	RawResponse []byte
	Errors      []Error `json:"Errors"`
	Status      string  `json:"status"`
	Message     string  `json:"message"`
}

// RemoveGallery removes a gallery and all of its subjects.
func (k *Kairos) RemoveGallery(galleryName string) (*ResponseRemoveGallery, error) {
	if galleryName == "" {
		return nil, fmt.Errorf("galleryName: should be present")
	}

	p := map[string]interface{}{
		"gallery_name": galleryName,
	}

	resp, err := k.makeRequest("POST", "gallery/remove", p)
	if err != nil {
		return nil, err
	}

	re := &ResponseRemoveGallery{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
