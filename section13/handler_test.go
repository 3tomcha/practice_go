package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/fortune", nil)
	w := httptest.NewRecorder()
	fortuneHandler(w, r)
	var j map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &j); err != nil {
		t.Errorf("json decode err should be nil: %v", err)
		return
	}
	if j["fortune"] != "大吉" {
		t.Errorf("result should be 大吉, but %s", j["fortune"])
	}
}
