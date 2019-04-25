package heartbeat

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"encoding/json"
)

func TestHeartbeat(t *testing.T) {
	req := httptest.NewRequest("GET", "/heartbeat", nil)
	r := httptest.NewRecorder()
	handler := &Handler{}
	handler.ServeHTTP(r, req)
	if r.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", r.Code, http.StatusOK)
	}
	var m Message
	err := json.Unmarshal([]byte(r.Body.String()), &m)
	if err != nil {
		t.Error("handler return invalid data")
	}
	if len(m.UntilNow) == 0 {
		t.Error("handler return incomplete info")
	}
}
