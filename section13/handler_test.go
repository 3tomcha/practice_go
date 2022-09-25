package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerRequest(t *testing.T) {
	s := httptest.NewServer(initServer())
	r, err := http.Get(s.URL + "/fortune")
	if err != nil {
		t.Errorf("http get err should be nil: %v", err)
		return
	}
	defer r.Body.Close()
	var j map[string]string
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		t.Errorf("json decode err should be nil: %v", err)
		return
	}
	if j["fortune"] != "大吉" {
		t.Errorf("result should be 大吉, but %s", j["fortune"])
	}
}