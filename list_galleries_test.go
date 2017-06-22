package kairgo_test

import (
	"net/http"
	"testing"
)

func Test_ListGalleries(t *testing.T) {
	setup()
	defer teardown()

	handleFun := func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		defer r.Body.Close()

		err := makeResponse(w, "gallery/list_all.json")
		if err != nil {
			t.Error(err)
		}
	}
	mux.HandleFunc("/gallery/list_all", handleFun)

	responseListGalleries, err := client.ListGalleries()
	if err != nil {
		t.Error(err)
		return
	}

	responseStatus := responseListGalleries.Status
	if responseStatus != "Complete" {
		t.Errorf("Expected %s, but was %s", "Complete", responseStatus)
	}

	galleryIDsCount := len(responseListGalleries.GalleryIDs)
	if galleryIDsCount != 5 {
		t.Errorf("Expected, but actual is %d", galleryIDsCount)
	}
}
