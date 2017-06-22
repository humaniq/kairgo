package kairgo_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func handleFuncViewSubject(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		var responsePath string

		defer r.Body.Close()
		body, err := requestBody(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		switch {
		case body["gallery_name"] != "MyGallery":
			responsePath = "gallery/view_subject_wrong_gallery.json"
		case body["subject_id"] != "test1":
			responsePath = "gallery/view_subject_wrong_subject.json"
		default:
			responsePath = "gallery/view_subject.json"
		}
		err = makeResponse(w, responsePath)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_ViewSubject(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/gallery/view_subject", handleFuncViewSubject(t))

	tests := []struct {
		galleryName  string
		subjectID    string
		status       string
		facesCount   int
		errorsCount  int
		errCode      int
		errorMessage string
	}{
		{
			galleryName:  "MyGallery",
			subjectID:    "test1",
			status:       "Complete",
			facesCount:   5,
			errorsCount:  0,
			errCode:      0,
			errorMessage: "",
		},
		{
			galleryName:  "OtherGallery",
			subjectID:    "test1",
			status:       "",
			facesCount:   0,
			errorsCount:  1,
			errCode:      5004,
			errorMessage: "gallery name not found",
		},
		{
			galleryName:  "MyGallery",
			subjectID:    "test2",
			status:       "",
			facesCount:   0,
			errorsCount:  1,
			errCode:      5003,
			errorMessage: "subject ID was not found",
		},
	}

	for _, test := range tests {
		result, err := client.ViewSubject(test.galleryName, test.subjectID)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, test.status, result.Status)
		assert.Equal(t, test.errorsCount, len(result.Errors))
		assert.Equal(t, test.facesCount, len(result.Faces))
		if test.errorsCount > 0 {
			assert.Equal(t, test.errCode, result.Errors[0].ErrCode)
			assert.Equal(t, test.errorMessage, result.Errors[0].Message)
		}
	}

}
