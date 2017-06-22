package kairgo

import (
	"encoding/json"
	"fmt"
)

// DetectRequest
type DetectRequest struct {
	// Required Parameters

	// Publicly accessible URL or Base64 encoded photo.
	Image string

	// Optional Parameters

	// MinHeadScale defined by you.
	// Is used to set the ratio of the smallest face we should look for in the photo.
	// Accepts a value between .015 (1:64 scale) and .5 (1:2 scale). By default it is set at .015 (1:64 scale) if not specified.
	MinHeadScale float32

	// Selector used to adjust the face detector.
	// If not specified the default of FRONTAL is used.
	// Note that these optional parameters are not reliable for face recognition, but may be useful for face detection uses.
	Selector string
}

func (d *DetectRequest) IsValid() (bool, error) {
	if d.Image == "" {
		return false, fmt.Errorf("Image: shuld be required")
	}
	return true, nil
}

type ResponseDetect struct {
	RawResponse []byte
	Errors      []Error `json:"Errors"`
	Images      []struct {
		Status string `json:"status"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		File   string `json:"file"`
		Faces  []struct {
			Face_id         int     `json:"face_id"`
			Quality         float32 `json:"quality"`
			Roll            int     `json:"roll"`
			TopLeftX        int     `json:"topLeftX"`
			TopLeftY        int     `json:"topLeftY"`
			ChinTipX        int     `json:"chinTipX"`
			RightEyeCenterX int     `json:"rightEyeCenterX"`
			Yaw             int     `json:"yaw"`
			ChinTipY        int     `json:"chinTipY"`
			Confidence      float32 `json:"confidence"`
			Height          int     `json:"height"`
			RightEyeCenterY int     `json:"rightEyeCenterY"`
			Width           int     `json:"width"`
			LeftEyeCenterY  int     `json:"leftEyeCenterY"`
			LeftEyeCenterX  int     `json:"leftEyeCenterX"`
			Pitch           int     `json:"pitch"`
			Attributes      struct {
				Lips   string  `json:"lips"`
				Asian  float32 `json:"asian"`
				Gender struct {
					Type string `json:"type"`
				} `json:"gender"`
				Age      int     `json:"age"`
				Hispanic float32 `json:"hispanic"`
				Other    float32 `json:"other"`
				Black    float32 `json:"black"`
				White    float32 `json:"white"`
				Glasses  string  `json:"glasses"`
			} `json:"attributes"`
		} `json:"faces"`
	} `json:"images"`
}

// Detect takes a photo and returns the facial features it finds.
func (k *Kairos) Detect(detectRequest *DetectRequest) (*ResponseDetect, error) {
	_, err := detectRequest.IsValid()
	if err != nil {
		return nil, err
	}

	p := map[string]interface{}{
		"image": detectRequest.Image,
	}

	if detectRequest.Selector != "" {
		p["selector"] = detectRequest.Selector
	}

	if detectRequest.MinHeadScale > 0 {
		p["minHeadScale"] = detectRequest.MinHeadScale
	}

	resp, err := k.makeRequest("POST", "detect", p)
	if err != nil {
		return nil, err
	}

	re := &ResponseDetect{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
