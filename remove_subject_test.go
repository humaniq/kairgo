package kairgo_test

import (
	"encoding/json"
	"fmt"
	"github.com/humaniq/kairgo"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	removeSubjectExistName = "MyGallery"
	removeSubjectExistID   = "test1"
	removeSubjectFaceId    = "58f9034743ab64939482"
	removeSubjectSuccess   = `
{
    "status": "Complete",
    "message": "subject id test1 has been successfully removed"
}
`
	removeSubjectWithFaceSuccess = `
{
    "status": "Complete",
    "message": "subject id test1 with face id 58f9034743ab64939482 has been successfully removed"
}
`

	removeSubjectError = `
{
    "Errors": [
    {
        "Message": "gallery name not found",
        "ErrCode": 5004
    }]
}
`
)

func handleFuncRemoveSubject(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
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

		if successSubjectRequest(body) {
			_, ok := body["face_id"]
			if ok {
				fmt.Fprint(w, removeSubjectWithFaceSuccess)
			} else {
				fmt.Fprint(w, removeSubjectSuccess)
			}
		} else {

			fmt.Fprint(w, removeSubjectError)
		}
	}
}

func successSubjectRequest(body map[string]interface{}) bool {
	return (body["gallery_name"] == removeSubjectExistName &&
		body["subject_id"] == removeSubjectExistID)
}

func Test_RemoveSubject_Success(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/remove_subject", handleFuncRemoveSubject(t))

	responseSubject, err := client.RemoveSubject(&kairgo.RemoveSubjectRequest{
		SubjectID:   removeSubjectExistID,
		GalleryName: removeSubjectExistName,
	})
	if err != nil {
		t.Error(err)
		return
	}

	status := responseSubject.Status
	message := responseSubject.Message

	if status != "Complete" {
		t.Errorf("Expected '%s', but actual: '%s'", "Complete", status)
	}

	if message != "subject id test1 has been successfully removed" {
		t.Errorf("Expected '%s', but actual: '%s'", "subject id test1 has been successfully removed", message)
	}

	errorsCount := len(responseSubject.Errors)
	if errorsCount != 0 {
		t.Errorf("Expected %d, but actual: %d", 0, errorsCount)
	}
}

func Test_RemoveSubjectWithFaceID_Success(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/remove_subject", handleFuncRemoveSubject(t))

	responseSubject, err := client.RemoveSubject(&kairgo.RemoveSubjectRequest{
		SubjectID:   removeSubjectExistID,
		GalleryName: removeSubjectExistName,
		FaceID:      removeSubjectFaceId,
	})
	if err != nil {
		t.Error(err)
		return
	}

	status := responseSubject.Status
	message := responseSubject.Message

	if status != "Complete" {
		t.Errorf("Expected '%s', but actual: '%s'", "Complete", status)
	}

	if message != "subject id test1 with face id 58f9034743ab64939482 has been successfully removed" {
		t.Errorf("Expected '%s', but actual: '%s'", "subject id test1 with face id 58f9034743ab64939482 has been successfully removed", message)
	}

	errorsCount := len(responseSubject.Errors)
	if errorsCount != 0 {
		t.Errorf("Expected %d, but actual: %d", 0, errorsCount)
	}
}

func Test_RemoveSubject_Fail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/gallery/remove_subject", handleFuncRemoveSubject(t))

	responseSubject, err := client.RemoveSubject(&kairgo.RemoveSubjectRequest{
		SubjectID:   "WrongSubjectID",
		GalleryName: removeSubjectExistName,
	})
	if err != nil {
		t.Error(err)
		return
	}

	status := responseSubject.Status

	if status != "" {
		t.Errorf("Expected '%s', but actual: '%s'", "", status)
	}

	errorsCount := len(responseSubject.Errors)
	if errorsCount == 0 {
		t.Errorf("Expected %d, but actual: %d", 1, errorsCount)
	}
}
