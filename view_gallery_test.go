package kairgo_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	galleryExistsName    = "exist_gallery"
	galleryWrongName     = "wrong_name"
	viewGalleriesSuccess = `
	{
	  "status": "Complete",
	  "subject_ids": [
	    "Elizabeth",
	    "Rachel"
	    ]
	}`

	viewGalleriesError = `
	{
	    "Errors": [
	    {
		"Message": "gallery name not found",
		"ErrCode": 5004
	    }]
	}
`
)

func handleFunc(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		defer r.Body.Close()
		body := make(map[string]interface{})

		b, rErr := ioutil.ReadAll(r.Body)
		if rErr != nil {
			t.Error(rErr)
		}

		uErr := json.Unmarshal(b, &body)
		if uErr != nil {
			t.Error(rErr)
		}

		if body["gallery_name"] == galleryExistsName {
			fmt.Fprint(w, viewGalleriesSuccess)
		} else {

			fmt.Fprint(w, viewGalleriesError)
		}
	}

}

func Test_ViewGallery_Success(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/view", handleFunc(t))

	responseGallery, err := client.ViewGallery(galleryExistsName)
	if err != nil {
		t.Error(err)
	}

	status := responseGallery.Status

	if status != "Complete" {
		t.Errorf("Expected '%s', but actual: '%s'", "Complete", status)
	}

	errorsCount := len(responseGallery.Errors)
	if errorsCount != 0 {
		t.Errorf("Expected %d, but actual: %d", 0, errorsCount)
	}
}

func Test_ViewGallery_Fail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/view", handleFunc(t))

	responseGallery, err := client.ViewGallery(galleryWrongName)
	if err != nil {
		t.Error(err)
	}

	status := responseGallery.Status

	if status != "" {
		t.Errorf("Expected '%s', but actual: '%s'", "", status)
	}

	errorsCount := len(responseGallery.Errors)
	if errorsCount == 0 {
		t.Errorf("Expected %d, but actual: %d", 1, errorsCount)
	}
}
