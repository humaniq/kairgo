package kairgo

import "encoding/json"

// ResponseEnroll ...
type ResponseEnroll struct {
	RawResponse []byte
	Errors      []Error `json:"Errors"`
	FaceID      string  `json:"face_id"`
	Images      []struct {
		Attributes struct {
			Lips   string  `json:"lips"`
			Asian  float64 `json:"asian"`
			Gender struct {
				Type string `json:"type"`
			} `json:"gender"`
			Age      int     `json:"age"`
			Hispanic float64 `json:"hispanic"`
			Other    float64 `json:"other"`
			Black    float64 `json:"black"`
			White    float64 `json:"white"`
			Glasses  string  `json:"glasses"`
		} `json:"attributes"`
		Transaction struct {
			Status      string  `json:"status"`
			TopLeftX    int     `json:"topLeftX"`
			TopLeftY    int     `json:"topLeftY"`
			GalleryName string  `json:"gallery_name"`
			Timestamp   string  `json:"timestamp"`
			Height      int     `json:"height"`
			Quality     float64 `json:"quality"`
			Confidence  float64 `json:"confidence"`
			SubjectID   string  `json:"subject_id"`
			Width       int     `json:"width"`
			FaceID      int     `json:"face_id"`
		} `json:"transaction"`
	} `json:"images"`
}

// Enroll an image
func (k *Kairos) Enroll(image, subjectID, galleryName, minHeadScale string, multipleFaces bool) (*ResponseEnroll, error) {
	p := make(map[string]interface{})
	p["image"] = image
	p["subject_id"] = subjectID
	p["gallery_name"] = galleryName

	// optional parameters
	if minHeadScale != "" {
		p["minHeadScale"] = minHeadScale
	}

	if multipleFaces != false {
		p["multiple_faces"] = multipleFaces
	}

	b, mErr := json.Marshal(p)
	if mErr != nil {
		return nil, mErr
	}

	req, reqErr := k.newRequest("POST", "enroll", b)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, doErr := k.do(req)
	if doErr != nil {
		return nil, doErr
	}

	re := &ResponseEnroll{}
	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
