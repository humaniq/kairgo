package kairgo_test

import (
	"github.com/humaniq/kairgo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	removeSubjectExistName = "MyGallery"
	removeSubjectExistID   = "test1"
	removeSubjectFaceId    = "58f9034743ab64939482"
)

func handleFuncRemoveSubject(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		var responsePath string

		defer r.Body.Close()
		body, err := requestBody(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		if successSubjectRequest(body) {
			_, ok := body["face_id"]
			if ok {
				responsePath = "gallery/remove_subject_with_face.json"
			} else {
				responsePath = "gallery/remove_subject.json"
			}
		} else {
			responsePath = "gallery/remove_subject_error.json"
		}

		err = makeResponse(w, responsePath)
		if err != nil {
			t.Error(err)
		}
	}
}

func successSubjectRequest(body map[string]interface{}) bool {
	return (body["gallery_name"] == removeSubjectExistName &&
		body["subject_id"] == removeSubjectExistID)
}

func Test_RemoveSubject(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/remove_subject", handleFuncRemoveSubject(t))

	tests := []struct {
		removeSubjectRequest kairgo.RemoveSubjectRequest
		status               string
		errorsCount          int
		errCode              int
		errorMessage         string
	}{
		{
			removeSubjectRequest: kairgo.RemoveSubjectRequest{
				GalleryName: "MyGallery",
				SubjectID:   "test1",
			},
			status:       "Complete",
			errorsCount:  0,
			errCode:      0,
			errorMessage: "",
		},
		{
			removeSubjectRequest: kairgo.RemoveSubjectRequest{
				GalleryName: "MyGallery",
				SubjectID:   "test1",
				FaceID:      "58f9034743ab64939482",
			},
			status:       "Complete",
			errorsCount:  0,
			errCode:      0,
			errorMessage: "",
		},
		{
			removeSubjectRequest: kairgo.RemoveSubjectRequest{
				GalleryName: "MyGallery",
				SubjectID:   "test2",
			},
			status:       "",
			errorsCount:  1,
			errCode:      5004,
			errorMessage: "gallery name not found",
		},
	}

	for _, test := range tests {
		result, err := client.RemoveSubject(&test.removeSubjectRequest)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, test.status, result.Status)
		assert.Equal(t, test.errorsCount, len(result.Errors))
		if test.errorsCount > 0 {
			assert.Equal(t, test.errCode, result.Errors[0].ErrCode)
			assert.Equal(t, test.errorMessage, result.Errors[0].Message)
		}
	}
}
