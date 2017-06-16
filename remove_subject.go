package kairgo

import (
	"encoding/json"
	"fmt"
)

type RemoveSubjectRequest struct {
	SubjectID, GalleryName string //required fields
	FaceID                 string //optional fields
}

func (r *RemoveSubjectRequest) IsValid() (bool, error) {
	if r.SubjectID == "" {
		return false, fmt.Errorf("SubjectID: shuld be required")
	}

	if r.GalleryName == "" {
		return false, fmt.Errorf("GalleryName: shuld be required")
	}

	return true, nil
}

type ResponseRemoveSubject struct {
	RawResponse []byte
	Errors      []Error `json:"Errors"`
	Status      string  `json:"status"`
	Message     string  `json:"message"`
}

// RemoveSubject removes a face you have enrolled within a gallery.
func (k *Kairos) RemoveSubject(removeSubjectRequest *RemoveSubjectRequest) (*ResponseRemoveSubject, error) {
	_, err := removeSubjectRequest.IsValid()
	if err != nil {
		return nil, err
	}

	p := map[string]interface{}{
		"gallery_name": removeSubjectRequest.GalleryName,
		"subject_id":   removeSubjectRequest.SubjectID,
	}

	if removeSubjectRequest.FaceID != "" {
		p["face_id"] = removeSubjectRequest.FaceID
	}

	resp, err := k.makeRequest("POST", "gallery/remove_subject", p)
	if err != nil {
		return nil, err
	}

	re := &ResponseRemoveSubject{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return nil, uErr
	}

	re.RawResponse = resp
	return re, nil
}
