package kairgo

import "encoding/json"

// ResponseRecognize ...
type ResponseRecognize struct {
	Errors      []Error `json:"Errors"`
	RawResponse []byte
	Images      []struct {
		Transaction struct {
			Status      string  `json:"status"`
			Width       int     `json:"width"`
			TopLeftX    int     `json:"topLeftX"`
			TopLeftY    int     `json:"topLeftY"`
			GalleryName string  `json:"gallery_name"`
			FaceID      int     `json:"face_id"`
			Confidence  float64 `json:"confidence"`
			SubjectID   string  `json:"subject_id"`
			Height      int     `json:"height"`
			Quality     float64 `json:"quality"`
		} `json:"transaction"`
		Candidates []struct {
			SubjectID           string  `json:"subject_id"`
			Confidence          float64 `json:"confidence"`
			EnrollmentTimestamp string  `json:"enrollment_timestamp"`
		} `json:"candidates"`
	} `json:"images"`
}

// Recognize takes a photo, finds the faces within it,
// and tries to match them against the faces you have already enrolled into a gallery.
func (k *Kairos) Recognize(image, galleryName, minHeadScale, threshold string, maxNumResults int) (*ResponseRecognize, error) {
	p := make(map[string]interface{})
	p["image"] = image
	p["gallery_name"] = galleryName

	// optional parameters
	if minHeadScale != "" {
		p["minHeadScale"] = minHeadScale

	}

	if threshold != "" {
		p["threshold"] = threshold

	}

	if maxNumResults != 0 {

		p["max_num_results"] = maxNumResults
	}

	b, mErr := json.Marshal(p)
	if mErr != nil {
		return nil, mErr
	}

	req, reqErr := k.newRequest("POST", "recognize", b)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, doErr := k.do(req)
	if doErr != nil {
		return nil, doErr
	}

	re := &ResponseRecognize{}
	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
