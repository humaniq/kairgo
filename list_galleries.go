package kairgo

import "encoding/json"

type ResponseListGalleries struct {
	RawResponse []byte
	Status      string   `json:"status"`
	GalleryIDs  []string `json:"gallery_ids"`
}

// ListGalleries ...
func (k *Kairos) ListGalleries() (*ResponseListGalleries, error) {
	req, reqErr := k.newRequest("POST", "gallery/list_all", nil)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, doErr := k.do(req)
	if doErr != nil {
		return nil, doErr
	}

	re := &ResponseListGalleries{}
	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
