package kairgo

import (
	"encoding/json"
	"fmt"
)

type ResponseVerify struct {
	RawResponse []byte
	Errors      []Error `json:"Errors"`
	Images      []struct {
		Transaction struct {
			Status      string  `json:"status"`
			Subject_id  string  `json:"subject_id"`
			Quality     float32 `json:"quality"`
			Width       int     `json:"width"`
			Height      int     `json:"height"`
			TopLeftX    int     `json:"topLeftX"`
			TopLeftY    int     `json:"topLeftY"`
			Confidence  float32 `json:"confidence"`
			GalleryName string  `json:"gallery_name"`
		} `json:"transaction"`
	} `json:"images"`
}

func (k *Kairos) Verify(image, galleryName, subjectID string) (*ResponseVerify, error) {
	if (image == "") || (galleryName == "") || (subjectID == "") {
		return nil, fmt.Errorf("All methods are required")
	}

	p := map[string]interface{}{
		"image":        image,
		"gallery_name": galleryName,
		"subject_id":   subjectID,
	}

	resp, err := k.makeRequest("POST", "verify", p)
	if err != nil {
		return nil, err
	}

	re := &ResponseVerify{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
