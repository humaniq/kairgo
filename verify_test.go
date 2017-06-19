package kairgo_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	verifyImageSuccess       = "http://media.kairos.com/kairos-elizabeth2.jpg"
	verifyGalleryNameSuccess = "MyGallery"
	verifySubjectID          = "Elizabeth"
)

func handleFuncVerify(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
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
		case body["image"] != verifyImageSuccess:
			responsePath = "verify_wrong_image.json"
		case body["gallery_name"] != verifyGalleryNameSuccess:
			responsePath = "verify_wrong_gallery.json"
		case body["subject_id"] != verifySubjectID:
			responsePath = "verify_wrong_subject.json"
		default:
			responsePath = "verify.json"
		}

		err = makeResponse(w, responsePath)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_Verify(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/verify", handleFuncVerify(t))

	tests := []struct {
		image, galleryName, subjectID string
		imagesCount                   int
		errorsCount                   int
		errCode                       int
		errorMessage                  string
	}{
		{
			image:        verifyImageSuccess,
			galleryName:  verifyGalleryNameSuccess,
			subjectID:    verifySubjectID,
			imagesCount:  1,
			errorsCount:  0,
			errCode:      0,
			errorMessage: "",
		}, {

			image:        "wrongImage",
			galleryName:  verifyGalleryNameSuccess,
			subjectID:    verifySubjectID,
			imagesCount:  0,
			errorsCount:  1,
			errCode:      5002,
			errorMessage: "no faces found in the image",
		}, {

			image:        verifyImageSuccess,
			galleryName:  "wongGallery",
			subjectID:    verifySubjectID,
			imagesCount:  0,
			errorsCount:  1,
			errCode:      5004,
			errorMessage: "gallery name not found",
		}, {

			image:        verifyImageSuccess,
			galleryName:  verifyGalleryNameSuccess,
			subjectID:    "wrongSubject",
			imagesCount:  0,
			errorsCount:  1,
			errCode:      5003,
			errorMessage: "subject ID was not found",
		},
	}
	for _, test := range tests {
		result, err := client.Verify(test.image, test.galleryName, test.subjectID)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, test.errorsCount, len(result.Errors))
		assert.Equal(t, test.imagesCount, len(result.Images))
		if test.errorsCount > 0 {
			assert.Equal(t, test.errCode, result.Errors[0].ErrCode)
			assert.Equal(t, test.errorMessage, result.Errors[0].Message)
		}
	}
}
