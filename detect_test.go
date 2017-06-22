package kairgo_test

import (
	"github.com/humaniq/kairgo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	detectImageUrlSuccess = "http://media.kairos.com/kairos-elizabeth.jpg"
	detectImageUrlWrong   = "http://media.kairos.com/kairos-elizabeth.txt"
)

func handleFuncDetect(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		var responsePath string

		defer r.Body.Close()

		body, err := requestBody(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		if body["image"] == detectImageUrlSuccess {
			responsePath = "detect.json"
		} else {
			responsePath = "detect_wrong.json"

		}

		err = makeResponse(w, responsePath)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_Detect(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/detect", handleFuncDetect(t))

	tests := []struct {
		detectRequest kairgo.DetectRequest
		status        string
		errorsCount   int
		errCode       int
		errorMessage  string
	}{
		{
			detectRequest: kairgo.DetectRequest{Image: detectImageUrlSuccess},
			status:        "Complete",
			errorsCount:   0,
			errCode:       0,
			errorMessage:  "",
		},
		{
			detectRequest: kairgo.DetectRequest{Image: detectImageUrlWrong},
			status:        "",
			errorsCount:   1,
			errCode:       5002,
			errorMessage:  "no faces found in the image",
		},
	}

	for _, test := range tests {
		result, err := client.Detect(&test.detectRequest)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, test.errorsCount, len(result.Errors))
		if test.errorsCount > 0 {
			assert.Equal(t, test.errCode, result.Errors[0].ErrCode)
			assert.Equal(t, test.errorMessage, result.Errors[0].Message)
		} else {
			assert.Equal(t, test.status, result.Images[0].Status)
		}
	}
}
