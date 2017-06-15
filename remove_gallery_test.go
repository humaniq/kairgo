package kairgo_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	removeGalleryExistName = "MyGallery"
	removeGalleryWrongName = "wrong_name"
	removeGalleriesSuccess = `
	{
	    "status":"Complete",
	    "message":"gallery MyGallery was removed"
	}`

	removeGalleriesError = `
	{
	    "Errors": [
	    {
		"Message": "gallery name not found",
		"ErrCode": 5004
	    }]
	}
`
)

func handleFuncRmGallery(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
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

		if body["gallery_name"] == removeGalleryExistName {
			fmt.Fprint(w, removeGalleriesSuccess)
		} else {

			fmt.Fprint(w, removeGalleriesError)
		}
	}

}

func Test_RemoveGallery_Success(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/remove", handleFuncRmGallery(t))

	responseGallery, err := client.RemoveGallery(removeGalleryExistName)
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

func Test_RemoveGallery_Fail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/remove", handleFuncRmGallery(t))

	responseGallery, err := client.RemoveGallery(removeGalleryWrongName)
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
