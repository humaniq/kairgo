package kairgo

import (
	"encoding/json"
	"fmt"
)

type ResponseSubject struct {
	RawResponse []byte
	Errors      []Error `json:"Errors"`
	Status      string  `json:"status"`
	Faces       []struct {
		FaceID       string `json:"face_id"`
		EnrollmentTS string `json:"enrollment_timestamp"`
	} `json:"message"`
}

// ViewSubject ...
func (k *Kairos) ViewSubject(galleryName, subjectId string) (*ResponseSubject, error) {
	if galleryName == "" {
		return nil, fmt.Errorf("%s: should be required", "galleryName")
	}

	if subjectId == "" {
		return nil, fmt.Errorf("%s: should be required", "subjectId")
	}

	p := map[string]interface{}{
		"gallery_name": galleryName,
		"subject_id":   subjectId,
	}

	resp, err := k.makeRequest("POST", "gallery/view_subject", p)
	if err != nil {
		return nil, err
	}

	re := &ResponseSubject{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil

}
