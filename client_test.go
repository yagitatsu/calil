package calil

import (
	"testing"
	"net/http"
)

const dummyAppKey = "2bc265ea827cb23b11d1ee80a25ef575"


func TestSearchLibrary(t *testing.T) {
	client := NewClient(dummyAppKey, http.DefaultClient)
	result, err := client.SearchLibrary(SearchLibraryParams{
		Pref:   "東京都",
		City:   "世田谷区",
		Format: "json",
		Limit:  10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	t.Logf("result=%v", result)

	if result.Libraries[0].City != "世田谷区" {
		t.Errorf("city is expected: `世田谷区`, actual: `%s`", result.Libraries[0].City)
	}
}
