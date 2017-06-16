package kairgo_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func handleViewFunc(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		var responsePath string

		defer r.Body.Close()
		body, err := requestBody(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		if body["gallery_name"] == "MyGallery" {
			responsePath = "gallery/view.json"
		} else {
			responsePath = "gallery/view_error.json"
		}

		err = makeResponse(w, responsePath)
		if err != nil {
			t.Error(err)
		}
	}

}

func Test_ViewGallery(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/view", handleViewFunc(t))

	tests := []struct {
		galleryName  string
		subjectID    string
		status       string
		errorsCount  int
		errCode      int
		errorMessage string
	}{
		{
			galleryName:  "MyGallery",
			status:       "Complete",
			errorsCount:  0,
			errCode:      0,
			errorMessage: "",
		},
		{
			galleryName:  "WrongName",
			status:       "",
			errorsCount:  1,
			errCode:      5004,
			errorMessage: "gallery name not found",
		},
	}

	for _, test := range tests {
		result, err := client.ViewGallery(test.galleryName)
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
